package expression

import (
	"fmt"
	"github.com/opspec-io/sdk-golang/data"
	"github.com/opspec-io/sdk-golang/expression/interpolater"
	"github.com/opspec-io/sdk-golang/model"
)

type evalArrayer interface {
	// EvalToArray evaluates an expression to an array value
	// expression must be a type supported by data.CoerceToArray
	EvalToArray(
		scope map[string]*model.Value,
		expression interface{},
		pkgHandle model.PkgHandle,
	) (*model.Value, error)
}

func newEvalArrayer() evalArrayer {
	return _evalArrayer{
		evalArrayInitializerer: newEvalArrayInitializerer(),
		data:         data.New(),
		interpolater: interpolater.New(),
	}
}

type _evalArrayer struct {
	evalArrayInitializerer
	data         data.Data
	interpolater interpolater.Interpolater
}

func (ea _evalArrayer) EvalToArray(
	scope map[string]*model.Value,
	expression interface{},
	pkgHandle model.PkgHandle,
) (*model.Value, error) {
	switch expression := expression.(type) {
	case []interface{}:
		arrayValue, err := ea.evalArrayInitializerer.Eval(
			expression,
			scope,
			pkgHandle,
		)
		if nil != err {
			return nil, fmt.Errorf("unable to evaluate %+v to array; error was %v", expression, err)
		}

		return &model.Value{Array: arrayValue}, nil
	case string:
		var value *model.Value
		if ref, ok := tryResolveExplicitRef(expression, scope); ok {
			value = ref
		} else {
			stringValue, err := ea.interpolater.Interpolate(
				expression,
				scope,
				pkgHandle,
			)
			if nil != err {
				return nil, err
			}
			value = &model.Value{String: &stringValue}
		}
		return ea.data.CoerceToArray(value)
	}
	return nil, fmt.Errorf("unable to evaluate %+v to array; unsupported type", expression)
}
