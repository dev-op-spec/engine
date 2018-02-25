package inputs

//go:generate counterfeiter -o ./fake.go --fake-name Fake ./ Inputs

import (
	"github.com/opspec-io/sdk-golang/data"
	"github.com/opspec-io/sdk-golang/model"
)

type Inputs interface {
	// Validate validates inputVals against inputParams
	// note: param defaults aren't considered
	Validate(
		inputVals map[string]*model.Value,
		inputParams map[string]*model.Param,
	) map[string][]error

	// Interpret interprets inputs via the provided inputArgs, inputParams, and scope;
	// opScratchDir will be used to store any run data such as scope coercions to files
	Interpret(
		inputArgs map[string]interface{},
		inputParams map[string]*model.Param,
		parentPkgHandle model.PkgHandle,
		pkgPath string,
		scope map[string]*model.Value,
		opScratchDir string,
	) (map[string]*model.Value, []error)
}

func New() Inputs {
	return _Inputs{
		argInterpreter: newArgInterpreter(),
		data:           data.New(),
	}
}

type _Inputs struct {
	argInterpreter argInterpreter
	data           data.Data
}
