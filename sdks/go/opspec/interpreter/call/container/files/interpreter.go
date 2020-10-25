package files

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/golang-interfaces/ios"
	"github.com/golang-utils/filecopier"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/file"
)

//counterfeiter:generate -o fakes/interpreter.go . Interpreter
type Interpreter interface {
	Interpret(
		scope map[string]*model.Value,
		containerCallSpecFiles map[string]interface{},
		scratchDirPath string,
	) (map[string]string, error)
}

// NewInterpreter returns an initialized Interpreter instance
func NewInterpreter(
	dataDirPath string,
) Interpreter {
	return _interpreter{
		fileCopier:      filecopier.New(),
		fileInterpreter: file.NewInterpreter(),
		os:              ios.New(),
		dataDirPath:     dataDirPath,
	}
}

type _interpreter struct {
	fileCopier      filecopier.FileCopier
	fileInterpreter file.Interpreter
	os              ios.IOS
	dataDirPath     string
}

func (itp _interpreter) Interpret(
	scope map[string]*model.Value,
	containerCallSpecFiles map[string]interface{},
	scratchDirPath string,
) (map[string]string, error) {
	containerCallFiles := map[string]string{}
fileLoop:
	for callSpecContainerFilePath, fileExpression := range containerCallSpecFiles {

		if nil == fileExpression {
			// bound implicitly
			fileExpression = fmt.Sprintf("$(%v)", callSpecContainerFilePath)
		}

		fileValue, err := itp.fileInterpreter.Interpret(
			scope,
			fileExpression,
			scratchDirPath,
			true,
		)
		if nil != err {
			return nil, fmt.Errorf(
				"unable to bind %v to %v; error was %v",
				callSpecContainerFilePath,
				fileExpression,
				err,
			)
		}

		if !strings.HasPrefix(*fileValue.File, itp.dataDirPath) {
			// bound to non rootFS file
			containerCallFiles[callSpecContainerFilePath] = *fileValue.File
			continue fileLoop
		}
		containerCallFiles[callSpecContainerFilePath] = filepath.Join(scratchDirPath, callSpecContainerFilePath)

		// create file
		if err := itp.os.MkdirAll(
			filepath.Dir(containerCallFiles[callSpecContainerFilePath]),
			0777,
		); nil != err {
			return nil, fmt.Errorf(
				"unable to bind %v to %v; error was %v",
				callSpecContainerFilePath,
				fileExpression,
				err,
			)
		}

		err = itp.fileCopier.OS(
			*fileValue.File,
			containerCallFiles[callSpecContainerFilePath],
		)
		if nil != err {
			return nil, fmt.Errorf(
				"unable to bind %v to %v; error was %v",
				callSpecContainerFilePath,
				fileExpression,
				err,
			)
		}

	}
	return containerCallFiles, nil
}
