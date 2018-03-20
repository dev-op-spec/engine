package expression

import (
	"fmt"
	"github.com/golang-interfaces/encoding-ijson"
	"github.com/opspec-io/sdk-golang/data/coerce"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/op/interpreter/expression/interpolater"
)

//go:generate counterfeiter -o ./fakeEvalObjectInitializerer.go --fake-name fakeEvalObjectInitializerer ./ evalObjectInitializerer

type evalObjectInitializerer interface {
	// Eval evaluates an objectInitializer expression
	Eval(
		expression map[string]interface{},
		scope map[string]*model.Value,
		opHandle model.DataHandle,
	) (map[string]interface{}, error)
}

// newEvalObjectInitializerer returns a new evalObjectInitializerer instance
func newEvalObjectInitializerer() evalObjectInitializerer {
	return _evalObjectInitializerer{
		coerce:       coerce.New(),
		interpolater: interpolater.New(),
		json:         ijson.New(),
	}
}

type _evalObjectInitializerer struct {
	coerce       coerce.Coerce
	interpolater interpolater.Interpolater
	json         ijson.IJSON
}

func (eoi _evalObjectInitializerer) Eval(
	expression map[string]interface{},
	scope map[string]*model.Value,
	opHandle model.DataHandle,
) (map[string]interface{}, error) {

	// expand shorthand properties w/out mutating original (maps passed by reference in go)
	expressionWithExpandedShorthandProps := map[string]interface{}{}
	for propName, propValue := range expression {
		if nil == propValue {
			expressionWithExpandedShorthandProps[propName] = fmt.Sprintf("%v%v%v", interpolater.RefStart, propName, interpolater.RefEnd)
		} else {
			expressionWithExpandedShorthandProps[propName] = propValue
		}
	}

	objectBytes, err := eoi.json.Marshal(expressionWithExpandedShorthandProps)
	if nil != err {
		return nil, fmt.Errorf("unable to eval %+v as objectInitializer; error was %v", expression, err)
	}

	objectJSON, err := eoi.interpolater.Interpolate(
		string(objectBytes),
		scope,
		opHandle,
	)
	if nil != err {
		return nil, fmt.Errorf("unable to eval %+v as objectInitializer; error was %v", expression, err)
	}

	object := map[string]interface{}{}
	if err := eoi.json.Unmarshal([]byte(objectJSON), &object); nil != err {
		return nil, fmt.Errorf("unable to eval %+v as objectInitializer; error was %v", expression, err)
	}

	return object, nil
}