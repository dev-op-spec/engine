package files

import (
	"fmt"
	"path/filepath"
	"strings"

	"os"

	"github.com/golang-utils/filecopier"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/file"
)

// Interpret container files
func Interpret(
	scope map[string]*model.Value,
	containerCallSpecFiles map[string]interface{},
	scratchDirPath string,
	dataCachePath string,
) (map[string]string, error) {
	containerCallFiles := map[string]string{}
fileLoop:
	for callSpecContainerFilePath, fileExpression := range containerCallSpecFiles {

		if nil == fileExpression {
			// bound implicitly
			fileExpression = fmt.Sprintf("$(%v)", callSpecContainerFilePath)
		}

		fileValue, err := file.Interpret(
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

		if !strings.HasPrefix(*fileValue.File, dataCachePath) {
			// bound to non dataCachePath
			containerCallFiles[callSpecContainerFilePath] = *fileValue.File
			continue fileLoop
		}

		// copy cached files to ensure can't be mutated
		containerCallFiles[callSpecContainerFilePath] = filepath.Join(scratchDirPath, callSpecContainerFilePath)

		// create parent dir
		if err := os.MkdirAll(
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

		// copy file
		fileCopier := filecopier.New()
		err = fileCopier.OS(
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
