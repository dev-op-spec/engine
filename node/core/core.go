// Package core defines the core interface for an opspec node
package core

import (
	"context"
	"path/filepath"

	"github.com/opctl/sdk-golang/data"
	"github.com/opctl/sdk-golang/model"
	"github.com/opctl/sdk-golang/node/core/containerruntime"
	"github.com/opctl/sdk-golang/opspec/interpreter/call"
	"github.com/opctl/sdk-golang/opspec/interpreter/call/container"
	"github.com/opctl/sdk-golang/opspec/interpreter/call/op"
	dotyml "github.com/opctl/sdk-golang/opspec/opfile"
	"github.com/opctl/sdk-golang/util/pubsub"
	"github.com/opctl/sdk-golang/util/uniquestring"
)

//go:generate counterfeiter -o ./fake.go --fake-name Fake ./ Core

type Core interface {
	GetEventStream(
		ctx context.Context,
		req *model.GetEventStreamReq,
	) (
		<-chan model.Event,
		<-chan error,
	)

	KillOp(
		req model.KillOpReq,
	)

	StartOp(
		ctx context.Context,
		req model.StartOpReq,
	) (
		callId string,
		err error,
	)

	// Resolve attempts to resolve an op via local filesystem or git
	// nil pullCreds will be ignored
	//
	// expected errs:
	//  - ErrDataProviderAuthentication on authentication failure
	//  - ErrDataProviderAuthorization on authorization failure
	//  - ErrDataRefResolution on resolution failure
	ResolveData(
		ctx context.Context,
		dataRef string,
		pullCreds *model.PullCreds,
	) (
		model.DataHandle,
		error,
	)
}

func New(
	pubSub pubsub.PubSub,
	containerRuntime containerruntime.ContainerRuntime,
	dataDirPath string,
) (core Core) {
	uniqueStringFactory := uniquestring.NewUniqueStringFactory()

	callStore := newCallStore()

	callKiller := newCallKiller(
		callStore,
		containerRuntime,
		pubSub,
	)

	opInterpreter := op.NewInterpreter(dataDirPath)

	caller := newCaller(
		call.NewInterpreter(
			container.NewInterpreter(dataDirPath),
			dataDirPath,
		),
		newContainerCaller(
			containerRuntime,
			pubSub,
		),
		dataDirPath,
		callStore,
		callKiller,
		pubSub,
	)

	core = _core{
		caller:              caller,
		containerRuntime:    containerRuntime,
		data:                data.New(),
		dataCachePath:       filepath.Join(dataDirPath, "pkgs"),
		callStore:           callStore,
		dotYmlGetter:        dotyml.NewGetter(),
		opInterpreter:       opInterpreter,
		callKiller:          callKiller,
		pubSub:              pubSub,
		uniqueStringFactory: uniqueStringFactory,
	}

	return
}

type _core struct {
	caller              caller
	containerRuntime    containerruntime.ContainerRuntime
	data                data.Data
	dataCachePath       string
	callStore           callStore
	dotYmlGetter        dotyml.Getter
	opInterpreter       op.Interpreter
	callKiller          callKiller
	pubSub              pubsub.PubSub
	uniqueStringFactory uniquestring.UniqueStringFactory
}
