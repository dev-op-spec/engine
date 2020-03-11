package eq

import (
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/str"
)

//counterfeiter:generate -o fakes/interpreter.go . Interpreter
type Interpreter interface {
	Interpret(
		expressions []interface{},
		opHandle model.DataHandle,
		scope map[string]*model.Value,
	) (bool, error)
}

// NewInterpreter returns an initialized Interpreter instance
func NewInterpreter() Interpreter {
	return &_interpreter{
		stringInterpreter: str.NewInterpreter(),
	}
}

type _interpreter struct {
	stringInterpreter str.Interpreter
}

func (itp _interpreter) Interpret(
	expressions []interface{},
	opHandle model.DataHandle,
	scope map[string]*model.Value,
) (bool, error) {
	var firstItemAsString string
	for i, expression := range expressions {
		// interpret items as strings since everything is coercible to string
		item, err := itp.stringInterpreter.Interpret(scope, expression, opHandle)
		if nil != err {
			return false, err
		}
		currentItemAsString := *item.String

		if i == 0 {
			firstItemAsString = currentItemAsString
			continue
		}

		if firstItemAsString != currentItemAsString {
			// if first seen item not equal to current item predicate is false.
			return false, nil
		}
	}
	return true, nil
}
