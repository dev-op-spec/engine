package paramvalidator

import (
	"errors"
	"fmt"
	"github.com/opspec-io/sdk-golang/model"
)

// validateFile validates an value against a file parameter
func (this _ParamValidator) validateFile(
	rawValue *string,
	param *model.FileParam,
) []error {

	value := rawValue
	if nil == value && nil != param.Default {
		// apply default if value not set
		value = param.Default
	}

	if nil == value {
		return []error{errors.New("File required")}
	}

	fileInfo, err := this.os.Stat(*value)
	if nil != err {
		return []error{err}
	} else if !fileInfo.Mode().IsRegular() {
		return []error{fmt.Errorf("%v not a file", *value)}
	}
	return []error{}
}
