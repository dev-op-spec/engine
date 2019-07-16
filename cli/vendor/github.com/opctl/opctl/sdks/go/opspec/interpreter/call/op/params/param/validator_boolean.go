package param

import (
	"errors"
	"github.com/opctl/opctl/sdks/go/model"
)

// validateBoolean validates a value against a boolean parameter
func (vdt _validator) validateBoolean(
	value *model.Value,
) []error {

	// handle no value passed
	if nil == value || nil == value.Boolean {
		return []error{errors.New("boolean required")}
	}

	return []error{}
}
