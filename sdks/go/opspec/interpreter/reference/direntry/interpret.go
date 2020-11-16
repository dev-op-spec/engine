package direntry

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/opctl/opctl/sdks/go/model"
)

// Interpret a dir entry ref i.e. refs of the form name/sub/file.ext
// it's an error if ref doesn't start with '/'
// returns ref remainder, dereferenced data, and error if one occurred
func Interpret(
	ref string,
	data *model.Value,
	opts *string,
) (string, *model.Value, error) {

	if !strings.HasPrefix(ref, "/") {
		return "", nil, fmt.Errorf("unable to interpret '%v' as dir entry ref; expected '/'", ref)
	}

	valuePath := filepath.Join(*data.Dir, ref)

	fileInfo, err := os.Stat(valuePath)
	if nil == err {
		if fileInfo.IsDir() {
			return "", &model.Value{Dir: &valuePath}, nil
		}

		return "", &model.Value{File: &valuePath}, nil
	} else if nil != opts && os.IsNotExist(err) {

		if "Dir" == *opts {
			err := os.MkdirAll(valuePath, 0700)
			if nil != err {
				return "", nil, fmt.Errorf("unable to interpret '%v' as dir entry ref; error was %v", ref, err)
			}

			return "", &model.Value{Dir: &valuePath}, nil
		}

		// handle file ref
		err := os.MkdirAll(filepath.Dir(valuePath), 0700)
		if nil != err {
			return "", nil, fmt.Errorf("unable to interpret '%v' as dir entry ref; error was %v", ref, err)
		}

		file, err := os.Create(valuePath)
		file.Close()
		if nil != err {
			return "", nil, fmt.Errorf("unable to interpret '%v' as dir entry ref; error was %v", ref, err)
		}

		return "", &model.Value{File: &valuePath}, nil

	}

	return "", nil, fmt.Errorf("unable to interpret '%v' as dir entry ref; error was %v", ref, err)

}
