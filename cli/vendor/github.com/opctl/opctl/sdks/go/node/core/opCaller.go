package core

import (
	"context"
	"errors"
	"path/filepath"
	"time"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/op/outputs"
	"github.com/opctl/opctl/sdks/go/opspec/opfile"
	"github.com/opctl/opctl/sdks/go/pubsub"
)

//counterfeiter:generate -o internal/fakes/opCaller.go . opCaller
type opCaller interface {
	// Executes an op call
	Call(
		ctx context.Context,
		opCall *model.OpCall,
		inboundScope map[string]*model.Value,
		parentCallID *string,
		opCallSpec *model.OpCallSpec,
	)
}

func newOpCaller(
	callStore callStore,
	pubSub pubsub.PubSub,
	caller caller,
	dataDirPath string,
) opCaller {
	return _opCaller{
		caller:             caller,
		callStore:          callStore,
		dcgScratchDir:      filepath.Join(dataDirPath, "dcg"),
		outputsInterpreter: outputs.NewInterpreter(),
		opFileGetter:       opfile.NewGetter(),
		pubSub:             pubSub,
	}
}

type _opCaller struct {
	callStore          callStore
	caller             caller
	dcgScratchDir      string
	outputsInterpreter outputs.Interpreter
	opFileGetter       opfile.Getter
	pubSub             pubsub.PubSub
}

func (oc _opCaller) Call(
	ctx context.Context,
	opCall *model.OpCall,
	inboundScope map[string]*model.Value,
	parentCallID *string,
	opCallSpec *model.OpCallSpec,
) {
	var err error
	outboundScope := map[string]*model.Value{}

	defer func() {
		// defer must be defined before conditional return statements so it always runs
		var opOutcome string
		if dcg := oc.callStore.TryGet(opCall.OpID); nil != dcg && dcg.IsKilled {
			opOutcome = model.OpOutcomeKilled
		} else if nil != err {
			opOutcome = model.OpOutcomeFailed
		} else {
			opOutcome = model.OpOutcomeSucceeded
		}

		event := model.Event{
			Timestamp: time.Now().UTC(),
			OpEnded: &model.OpEnded{
				OpID:     opCall.OpID,
				OpRef:    opCall.OpPath,
				Outcome:  opOutcome,
				RootOpID: opCall.RootOpID,
				Outputs:  outboundScope,
			},
		}

		if nil != err {
			event.OpEnded.Error = &model.CallEndedError{
				Message: err.Error(),
			}
		}

		oc.pubSub.Publish(event)

	}()

	opStartedTime := time.Now().UTC()

	oc.pubSub.Publish(
		model.Event{
			Timestamp: opStartedTime,
			OpStarted: &model.OpStarted{
				OpID:     opCall.OpID,
				OpRef:    opCall.OpPath,
				RootOpID: opCall.RootOpID,
			},
		},
	)

	// form scope for op call by combining defined inputs & op dir
	opCallScope := map[string]*model.Value{}
	for varName, varData := range opCall.Inputs {
		opCallScope[varName] = varData
	}
	opCallScope["/"] = &model.Value{
		Dir: &opCall.OpPath,
	}

	oc.caller.Call(
		ctx,
		opCall.ChildCallID,
		opCallScope,
		opCall.ChildCallCallSpec,
		opCall.OpPath,
		&opCall.OpID,
		opCall.RootOpID,
	)

	// subscribe to events
	eventChannel, _ := oc.pubSub.Subscribe(
		ctx,
		model.EventFilter{
			Roots: []string{opCall.RootOpID},
			Since: &opStartedTime,
		},
	)

	opOutputs := map[string]*model.Value{}

eventLoop:
	for event := range eventChannel {
		switch {
		case nil != event.OpEnded && event.OpEnded.OpID == opCall.OpID:
			// parent ended prematurely
			return
		case nil != event.CallEnded && event.CallEnded.CallID == opCall.ChildCallID:
			if nil != event.CallEnded.Error {
				err = errors.New(event.CallEnded.Error.Message)
				// end on any error
				return
			}
			for name, value := range event.CallEnded.Outputs {
				opOutputs[name] = value
			}
			break eventLoop
		}
	}

	var opFile *model.OpFile
	opFile, err = oc.opFileGetter.Get(
		ctx,
		opCall.OpPath,
	)
	if nil != err {
		return
	}
	opOutputs, err = oc.outputsInterpreter.Interpret(
		opOutputs,
		opFile.Outputs,
		opCall.OpPath,
		filepath.Join(oc.dcgScratchDir, opCall.OpID),
	)

	// filter op outboundScope to bound call outboundScope
	for boundName, boundValue := range opCallSpec.Outputs {
		// return bound outboundScope
		if "" == boundValue {
			// implicit value
			boundValue = boundName
		}
		for opOutputName, opOutputValue := range opOutputs {
			if boundValue == opOutputName {
				outboundScope[boundName] = opOutputValue
			}
		}
	}
}
