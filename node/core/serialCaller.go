package core

//go:generate counterfeiter -o ./fakeSerialCaller.go --fake-name fakeSerialCaller ./ serialCaller

import (
	"context"
	"errors"
	"time"

	"github.com/opctl/sdk-golang/model"
	"github.com/opctl/sdk-golang/util/pubsub"
	"github.com/opctl/sdk-golang/util/uniquestring"
)

type serialCaller interface {
	// Executes a serial call
	Call(
		ctx context.Context,
		callID string,
		inboundScope map[string]*model.Value,
		rootOpID string,
		opHandle model.DataHandle,
		scgSerialCall []*model.SCG,
	)
}

func newSerialCaller(
	caller caller,
	pubSub pubsub.PubSub,
) serialCaller {

	return _serialCaller{
		caller:              caller,
		pubSub:              pubSub,
		uniqueStringFactory: uniquestring.NewUniqueStringFactory(),
	}

}

type _serialCaller struct {
	caller              caller
	pubSub              pubsub.PubSub
	uniqueStringFactory uniquestring.UniqueStringFactory
}

func (sc _serialCaller) Call(
	ctx context.Context,
	callID string,
	inboundScope map[string]*model.Value,
	rootOpID string,
	opHandle model.DataHandle,
	scgSerialCall []*model.SCG,
) {
	outputs := map[string]*model.Value{}
	for varName, varData := range inboundScope {
		outputs[varName] = varData
	}
	var err error

	defer func() {
		// defer must be defined before conditional return statements so it always runs
		event := model.Event{
			SerialCallEnded: &model.SerialCallEndedEvent{
				CallID:   callID,
				Outputs:  outputs,
				RootOpID: rootOpID,
			},
			Timestamp: time.Now().UTC(),
		}

		if nil != err {
			event.SerialCallEnded.Error = &model.CallEndedEventError{
				Message: err.Error(),
			}
		}
		sc.pubSub.Publish(
			event,
		)
	}()

	// subscribe to events
	// @TODO: handle err channel
	eventFilterSince := time.Now().UTC()
	eventChannel, _ := sc.pubSub.Subscribe(
		ctx,
		model.EventFilter{
			Roots: []string{rootOpID},
			Since: &eventFilterSince,
		},
	)

	for _, scgCall := range scgSerialCall {

		var childCallID string
		childCallID, err = sc.uniqueStringFactory.Construct()
		if nil != err {
			// end run immediately on any error
			return
		}

		if err = sc.caller.Call(
			ctx,
			childCallID,
			outputs,
			scgCall,
			opHandle,
			&callID,
			rootOpID,
		); nil != err {
			// end run immediately on any error
			return
		}

	eventLoop:
		for event := range eventChannel {
			// merge child outboundScope w/ outboundScope, child outboundScope having precedence
			switch {
			case nil != event.CallEnded && event.CallEnded.CallID == childCallID:
				if nil != event.CallEnded.Error {
					err = errors.New(event.CallEnded.Error.Message)
				}
				for name, value := range event.CallEnded.Outputs {
					outputs[name] = value
				}
				break eventLoop
			}
		}

	}

}
