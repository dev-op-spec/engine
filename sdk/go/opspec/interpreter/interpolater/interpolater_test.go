package interpolater

import (
	"context"
	"fmt"
	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/sdk-golang/data"
	"github.com/opctl/sdk-golang/model"
	"io/ioutil"
	"os"
	"path/filepath"
)

var _ = Context("Interpolate", func() {
	Describe("called for every scenario", func() {
		It("should return result fulfilling scenario", func() {
			rootPath := "testdata/interpolater/Interpolate"
			data := data.New()
			fsProvider := data.NewFSProvider()

			filepath.Walk(rootPath,
				func(path string, info os.FileInfo, err error) error {
					if info.IsDir() {
						scenariosDotYmlFilePath := filepath.Join(path, "scenarios.yml")
						if _, err := os.Stat(scenariosDotYmlFilePath); nil == err {
							/* arrange */
							scenariosDotYmlBytes, err := ioutil.ReadFile(scenariosDotYmlFilePath)
							if nil != err {
								panic(err)
							}

							scenarioDotYml := []struct {
								Name     string
								Template string
								Scope    map[string]*model.Value
								Expected string
							}{}
							if err := yaml.Unmarshal(scenariosDotYmlBytes, &scenarioDotYml); nil != err {
								panic(fmt.Errorf("error unmarshalling scenario.yml for %v; error was %v", path, err))
							}

							absPath, err := filepath.Abs(path)
							if nil != err {
								panic(fmt.Errorf("error getting absPath for %v; error was %v", path, err))
							}

							opHandle, err := data.Resolve(context.Background(), absPath, fsProvider)
							if nil != err {
								panic(fmt.Errorf("error getting opHandle for %v; error was %v", path, err))
							}

							for _, scenario := range scenarioDotYml {
								for name, value := range scenario.Scope {
									// make file refs absolute
									if nil != value.File {
										absFilePath := filepath.Join(absPath, *value.File)
										scenario.Scope[name] = &model.Value{File: &absFilePath}
									}
								}
								/* act */
								objectUnderTest := New()
								actualResult, actualErr := objectUnderTest.Interpolate(
									scenario.Template,
									scenario.Scope,
									opHandle,
								)

								/* assert */
								description := fmt.Sprintf("scenario:\n  path: '%v'\n  name: '%v'", path, scenario.Name)
								Expect(actualErr).To(BeNil(), description)
								Expect(actualResult).To(Equal(scenario.Expected), description)
							}
						}
					}
					return nil
				})
		})
	})
})
