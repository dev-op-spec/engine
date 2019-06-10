package core

//go:generate counterfeiter -o ./fakeLooper.go --fake-name fakeLooper ./ looper

import (
	"context"
	"time"

	"github.com/opctl/sdk-golang/opspec/interpreter/call/loop"
	"github.com/opctl/sdk-golang/opspec/interpreter/call/loop/iteration"
	"github.com/opctl/sdk-golang/opspec/interpreter/loopable"

	"github.com/opctl/sdk-golang/model"
	"github.com/opctl/sdk-golang/util/pubsub"
	"github.com/opctl/sdk-golang/util/uniquestring"
)

type looper interface {
	// Loop loops a call
	Loop(
		ctx context.Context,
		id string,
		inboundScope map[string]*model.Value,
		scg *model.SCG,
		opHandle model.DataHandle,
		parentCallID *string,
		rootOpID string,
	) error
}

func newLooper(
	caller caller,
	pubSub pubsub.PubSub,
) looper {
	return _looper{
		caller:              caller,
		iterationScoper:     iteration.NewScoper(),
		loopableInterpreter: loopable.NewInterpreter(),
		loopDeScoper:        loop.NewDeScoper(),
		pubSub:              pubSub,
		uniqueStringFactory: uniquestring.NewUniqueStringFactory(),
		loopInterpreter:     loop.NewInterpreter(),
	}
}

type _looper struct {
	caller              caller
	iterationScoper     iteration.Scoper
	loopableInterpreter loopable.Interpreter
	loopDeScoper        loop.DeScoper
	pubSub              pubsub.PubSub
	uniqueStringFactory uniquestring.UniqueStringFactory
	loopInterpreter     loop.Interpreter
}

func (lpr _looper) isLoopEnded(
	index int,
	loop *model.DCGLoop,
) bool {
	if nil != loop.Until && *loop.Until {
		// exit condition provided & met
		return true
	}

	if nil != loop.For {
		if nil != loop.For.Each.Array {
			return index == len(*loop.For.Each.Array)
		}
		if nil != loop.For.Each.Object {
			return index == len(*loop.For.Each.Object)
		}

		// empty array or object
		return true
	}

	return false
}

func (lpr _looper) Loop(
	ctx context.Context,
	id string,
	inboundScope map[string]*model.Value,
	scg *model.SCG,
	opHandle model.DataHandle,
	parentCallID *string,
	rootOpID string,
) error {
	outboundScope := map[string]*model.Value{}

	defer func() {
		// defer must be defined before conditional return statements so it always runs
		lpr.pubSub.Publish(
			model.Event{
				Timestamp: time.Now().UTC(),
				CallEnded: &model.CallEndedEvent{
					CallID:     id,
					RootCallID: rootOpID,
					Outputs:    outboundScope,
				},
			},
		)
	}()

	index := 0
	var err error
	outboundScope, err = lpr.iterationScoper.Scope(
		index,
		inboundScope,
		scg.Loop,
		opHandle,
	)
	if nil != err {
		return err
	}

	// copy scg.Loop & remove from scg since we're already looping
	scgLoop := scg.Loop
	scg.Loop = nil

	// interpret initial iteration of the loop
	dcgLoop, err := lpr.loopInterpreter.Interpret(
		opHandle,
		scgLoop,
		outboundScope,
	)
	if nil != err {
		return err
	}

	for !lpr.isLoopEnded(index, dcgLoop) {
		eventFilterSince := time.Now().UTC()

		callID, err := lpr.uniqueStringFactory.Construct()
		if nil != err {
			return err
		}

		err = lpr.caller.Call(
			ctx,
			callID,
			outboundScope,
			scg,
			opHandle,
			parentCallID,
			rootOpID,
		)
		if nil != err {
			// end looping on any error
			return err
		}

		// subscribe to events
		// @TODO: handle err channel
		eventChannel, _ := lpr.pubSub.Subscribe(
			ctx,
			model.EventFilter{
				Roots: []string{rootOpID},
				Since: &eventFilterSince,
			},
		)

	eventLoop:
		for event := range eventChannel {
			// merge child outboundScope w/ outboundScope, child outboundScope having precedence
			switch {
			case nil != event.CallEnded && event.CallEnded.CallID == callID:
				for name, value := range event.CallEnded.Outputs {
					outboundScope[name] = value
				}
				break eventLoop
			}
		}

		index++

		if lpr.isLoopEnded(index, dcgLoop) {
			break
		}

		outboundScope, err = lpr.iterationScoper.Scope(
			index,
			outboundScope,
			scgLoop,
			opHandle,
		)
		if nil != err {
			return err
		}

		// interpret next iteration of the loop
		dcgLoop, err = lpr.loopInterpreter.Interpret(
			opHandle,
			scgLoop,
			outboundScope,
		)
		if nil != err {
			return err
		}
	}

	lpr.loopDeScoper.DeScope(
		inboundScope,
		scgLoop,
		outboundScope,
	)

	return nil
}
