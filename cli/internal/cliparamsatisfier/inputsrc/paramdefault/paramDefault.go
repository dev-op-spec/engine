package paramdefault

import (
	"github.com/opctl/opctl/cli/internal/cliparamsatisfier/inputsrc"
	"github.com/opctl/opctl/sdks/go/model"
	"strings"
)

func New(
	inputs map[string]*model.Param,
) inputsrc.InputSrc {
	return paramDefaultInputSrc{
		inputs:      inputs,
		readHistory: map[string]struct{}{},
	}
}

// paramDefaultInputSrc implements InputSrc interface by sourcing inputs from input defaults
type paramDefaultInputSrc struct {
	inputs      map[string]*model.Param
	readHistory map[string]struct{} // tracks reads
}

func (this paramDefaultInputSrc) ReadString(
	inputName string,
) (*string, bool) {
	if _, ok := this.readHistory[inputName]; ok {
		// enforce read at most once.
		return nil, false
	}

	if inputValue, ok := this.inputs[inputName]; ok {
		// track read history
		this.readHistory[inputName] = struct{}{}

		switch {
		case nil != inputValue.Array && nil != inputValue.Array.Default:
			return nil, true
		case nil != inputValue.Boolean && nil != inputValue.Boolean.Default:
			return nil, true
		case nil != inputValue.Dir && nil != inputValue.Dir.Default:
			if strings.HasPrefix(*inputValue.Dir.Default, "/") {
				// defaulted to pkg dir
				return nil, true
			}
			return inputValue.Dir.Default, true
		case nil != inputValue.File && nil != inputValue.File.Default:
			if strings.HasPrefix(*inputValue.File.Default, "/") {
				// defaulted to pkg file
				return nil, true
			}
			return inputValue.File.Default, true
		case nil != inputValue.Number && nil != inputValue.Number.Default:
			return nil, true
		case nil != inputValue.Object && nil != inputValue.Object.Default:
			return nil, true
		case nil != inputValue.String && nil != inputValue.String.Default:
			return nil, true
		}
	}
	return nil, false
}
