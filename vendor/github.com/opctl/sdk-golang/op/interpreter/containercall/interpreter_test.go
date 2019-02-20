package containercall

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/golang-interfaces/ios"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/sdk-golang/data"
	"github.com/opctl/sdk-golang/model"
	"github.com/opctl/sdk-golang/op/interpreter/containercall/dirs"
	"github.com/opctl/sdk-golang/op/interpreter/containercall/envvars"
	"github.com/opctl/sdk-golang/op/interpreter/containercall/files"
	"github.com/opctl/sdk-golang/op/interpreter/containercall/image"
	"github.com/opctl/sdk-golang/op/interpreter/containercall/sockets"
	stringPkg "github.com/opctl/sdk-golang/op/interpreter/string"
)

var _ = Context("Interpreter", func() {
	Context("NewInterpreter", func() {
		It("shouldn't return nil", func() {
			/* arrange/act/assert */
			Expect(NewInterpreter("")).To(Not(BeNil()))
		})
	})
	Context("Interpret", func() {
		It("calls os.MkdirAll w/ expected scratchdir path & returns error", func() {
			/* arrange */
			dataDirPath := "/dummydataDirPath"
			providedContainerID := "dummyContainerID"
			providedRootOpID := "dummyRootOpID"

			expectedScratchDirPath := filepath.Join(
				dataDirPath,
				"dcg",
				providedRootOpID,
				"containers",
				providedContainerID,
				"fs",
			)
			expectedScratchDirMode := os.FileMode(0700)

			expectedError := errors.New("dummyError")

			fakeOS := new(ios.Fake)
			// error to trigger immediate return
			fakeOS.MkdirAllReturns(expectedError)

			objectUnderTest := _interpreter{
				os:          fakeOS,
				dataDirPath: dataDirPath,
			}

			/* act */
			_, actualError := objectUnderTest.Interpret(
				map[string]*model.Value{},
				&model.SCGContainerCall{},
				providedContainerID,
				providedRootOpID,
				new(data.FakeHandle),
			)

			/* assert */
			actualScratchDirPath, actualScratchDirMode := fakeOS.MkdirAllArgsForCall(0)
			Expect(actualScratchDirPath).To(Equal(expectedScratchDirPath))
			Expect(actualScratchDirMode).To(Equal(expectedScratchDirMode))
			Expect(actualError).To(Equal(expectedError))
		})

		Context("container.Cmd not empty", func() {
			It("should call stringInterpreter.Interpret w/ expected args for each container.Cmd entry", func() {
				/* arrange */
				providedString1 := "dummyString1"
				providedCurrentScope := map[string]*model.Value{
					"name1": {String: &providedString1},
				}
				providedOpHandle := new(data.FakeHandle)

				providedSCGContainerCall := &model.SCGContainerCall{
					Cmd: []interface{}{
						"dummy1",
						"dummy2",
					},
				}

				fakeStringInterpreter := new(stringPkg.FakeInterpreter)
				fakeStringInterpreter.InterpretReturns(&model.Value{String: new(string)}, nil)

				objectUnderTest := _interpreter{
					dirsInterpreter:    new(dirs.FakeInterpreter),
					envVarsInterpreter: new(envvars.FakeInterpreter),
					filesInterpreter:   new(files.FakeInterpreter),
					imageInterpreter:   new(image.FakeInterpreter),
					stringInterpreter:  fakeStringInterpreter,
					os:                 new(ios.Fake),
					socketsInterpreter: new(sockets.FakeInterpreter),
				}

				/* act */
				objectUnderTest.Interpret(
					providedCurrentScope,
					providedSCGContainerCall,
					"dummyContainerID",
					"dummyRootOpID",
					providedOpHandle,
				)

				/* assert */
				for expectedCmdIndex, expectedCmdEntry := range providedSCGContainerCall.Cmd {
					actualScope,
						actualCmdEntry,
						actualOpHandle := fakeStringInterpreter.InterpretArgsForCall(expectedCmdIndex)

					Expect(actualCmdEntry).To(Equal(expectedCmdEntry))
					Expect(actualScope).To(Equal(providedCurrentScope))
					Expect(actualOpHandle).To(Equal(providedOpHandle))
				}
			})
			It("should return expected dcg.Cmd", func() {
				/* arrange */
				expectedCmd := []string{
					"dummyCmdEntry1",
					"dummyCmdEntry2",
				}

				providedSCGContainerCall := &model.SCGContainerCall{
					Cmd: []interface{}{
						"dummy1",
						"dummy2",
					},
				}

				fakeStringInterpreter := new(stringPkg.FakeInterpreter)
				fakeStringInterpreter.InterpretReturnsOnCall(0, &model.Value{String: &expectedCmd[0]}, nil)
				fakeStringInterpreter.InterpretReturnsOnCall(1, &model.Value{String: &expectedCmd[1]}, nil)

				objectUnderTest := _interpreter{
					dirsInterpreter:    new(dirs.FakeInterpreter),
					envVarsInterpreter: new(envvars.FakeInterpreter),
					filesInterpreter:   new(files.FakeInterpreter),
					imageInterpreter:   new(image.FakeInterpreter),
					stringInterpreter:  fakeStringInterpreter,
					os:                 new(ios.Fake),
					socketsInterpreter: new(sockets.FakeInterpreter),
				}

				/* act */
				actualDCGContainerCall, _ := objectUnderTest.Interpret(
					map[string]*model.Value{},
					providedSCGContainerCall,
					"dummyContainerID",
					"dummyRootOpID",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualDCGContainerCall.Cmd).To(Equal(expectedCmd))
			})
		})

		It("should call dirs w/ expected args", func() {
			/* arrange */
			containerDirBind := "dummyContainerDirBind"

			providedScope := map[string]*model.Value{
				containerDirBind: {String: new(string)},
			}

			dirName := "dummyDirName"
			providedSCGContainerCall := &model.SCGContainerCall{
				Dirs: map[string]string{
					// implicitly bound
					dirName: "",
				},
			}

			provideddataDirPath := "dummydataDirPath"
			providedContainerID := "dummyContainerID"
			providedRootOpID := "dummyRootOpID"
			providedOpHandle := new(data.FakeHandle)

			expectedScratchDirPath := filepath.Join(
				provideddataDirPath,
				"dcg",
				providedRootOpID,
				"containers",
				providedContainerID,
				"fs",
			)

			fakeDirsInterpreter := new(dirs.FakeInterpreter)

			objectUnderTest := _interpreter{
				dirsInterpreter:    fakeDirsInterpreter,
				envVarsInterpreter: new(envvars.FakeInterpreter),
				filesInterpreter:   new(files.FakeInterpreter),
				imageInterpreter:   new(image.FakeInterpreter),
				os:                 new(ios.Fake),
				dataDirPath:        provideddataDirPath,
				socketsInterpreter: new(sockets.FakeInterpreter),
			}

			/* act */
			objectUnderTest.Interpret(
				providedScope,
				providedSCGContainerCall,
				providedContainerID,
				providedRootOpID,
				providedOpHandle,
			)

			/* assert */
			actualOpHandle, actualScope, actualScgContainerCallDirs, actualScratchDir := fakeDirsInterpreter.InterpretArgsForCall(0)
			Expect(actualOpHandle).To(Equal(providedOpHandle))
			Expect(actualScope).To(Equal(providedScope))
			Expect(actualScgContainerCallDirs).To(Equal(providedSCGContainerCall.Dirs))
			Expect(actualScratchDir).To(Equal(expectedScratchDirPath))
		})
		Context("dirs.Interpret errors", func() {
			It("should return expected error", func() {
				/* arrange */
				expectedErr := errors.New("dummyError")
				fakeDirsInterpreter := new(dirs.FakeInterpreter)
				fakeDirsInterpreter.InterpretReturns(nil, expectedErr)

				objectUnderTest := _interpreter{
					dirsInterpreter:    fakeDirsInterpreter,
					envVarsInterpreter: new(envvars.FakeInterpreter),
					imageInterpreter:   new(image.FakeInterpreter),
					os:                 new(ios.Fake),
				}

				/* act */
				_, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGContainerCall{},
					"dummyContainerID",
					"dummyRootOpID",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
		Context("dirs.Interpret doesn't error", func() {
			It("should return expected dcgContainerCall.Dirs", func() {
				/* arrange */
				expectedDCGContainerCallDirs := map[string]string{
					"dummyName": "dummyValue",
				}

				fakeDirsInterpreter := new(dirs.FakeInterpreter)
				fakeDirsInterpreter.InterpretReturns(expectedDCGContainerCallDirs, nil)

				objectUnderTest := _interpreter{
					dirsInterpreter:    fakeDirsInterpreter,
					envVarsInterpreter: new(envvars.FakeInterpreter),
					filesInterpreter:   new(files.FakeInterpreter),
					imageInterpreter:   new(image.FakeInterpreter),
					os:                 new(ios.Fake),
					socketsInterpreter: new(sockets.FakeInterpreter),
				}

				/* act */
				actualResult, _ := objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGContainerCall{},
					"dummyContainerID",
					"dummyRootOpID",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualResult.Dirs).To(Equal(expectedDCGContainerCallDirs))
			})
		})

		It("should call envVars w/ expected args", func() {
			/* arrange */
			containerFileBind := "dummyContainerFileBind"

			providedScope := map[string]*model.Value{
				containerFileBind: {String: new(string)},
			}

			envVarName := "dummyEnvVarName"
			providedSCGContainerCall := &model.SCGContainerCall{
				EnvVars: map[string]interface{}{
					// implicitly bound
					envVarName: "",
				},
			}

			providedOpHandle := new(data.FakeHandle)

			fakeEnvVarsInterpreter := new(envvars.FakeInterpreter)

			objectUnderTest := _interpreter{
				dirsInterpreter:    new(dirs.FakeInterpreter),
				envVarsInterpreter: fakeEnvVarsInterpreter,
				filesInterpreter:   new(files.FakeInterpreter),
				imageInterpreter:   new(image.FakeInterpreter),
				os:                 new(ios.Fake),
				socketsInterpreter: new(sockets.FakeInterpreter),
			}

			/* act */
			objectUnderTest.Interpret(
				providedScope,
				providedSCGContainerCall,
				"dummyContainerID",
				"dummyRootOpID",
				providedOpHandle,
			)

			/* assert */
			actualScope,
				actualScgContainerCallEnvVars,
				actualOpHandle := fakeEnvVarsInterpreter.InterpretArgsForCall(0)

			Expect(actualScope).To(Equal(providedScope))
			Expect(actualScgContainerCallEnvVars).To(Equal(providedSCGContainerCall.EnvVars))
			Expect(actualOpHandle).To(Equal(providedOpHandle))

		})
		Context("envVars.Interpret errors", func() {
			It("should return expected error", func() {
				/* arrange */
				expectedErr := errors.New("dummyError")
				fakeEnvVarsInterpreter := new(envvars.FakeInterpreter)
				fakeEnvVarsInterpreter.InterpretReturns(nil, expectedErr)

				objectUnderTest := _interpreter{
					dirsInterpreter:    new(dirs.FakeInterpreter),
					envVarsInterpreter: fakeEnvVarsInterpreter,
					os:                 new(ios.Fake),
				}

				/* act */
				_, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGContainerCall{},
					"dummyContainerID",
					"dummyRootOpID",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
		Context("envVars.Interpret doesn't error", func() {
			It("should return expected dcgContainerCall.EnvVars", func() {
				/* arrange */
				expectedDCGContainerCallEnvVars := map[string]string{
					"dummyName": "dummyValue",
				}

				fakeEnvVarsInterpreter := new(envvars.FakeInterpreter)
				fakeEnvVarsInterpreter.InterpretReturns(expectedDCGContainerCallEnvVars, nil)

				objectUnderTest := _interpreter{
					dirsInterpreter:    new(dirs.FakeInterpreter),
					envVarsInterpreter: fakeEnvVarsInterpreter,
					filesInterpreter:   new(files.FakeInterpreter),
					imageInterpreter:   new(image.FakeInterpreter),
					os:                 new(ios.Fake),
					socketsInterpreter: new(sockets.FakeInterpreter),
				}

				/* act */
				actualResult, _ := objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGContainerCall{},
					"dummyContainerID",
					"dummyRootOpID",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualResult.EnvVars).To(Equal(expectedDCGContainerCallEnvVars))
			})
		})

		It("should call files w/ expected args", func() {
			/* arrange */
			containerFileBind := "dummyContainerFileBind"

			providedScope := map[string]*model.Value{
				containerFileBind: {String: new(string)},
			}

			fileName := "dummyFileName"
			providedSCGContainerCall := &model.SCGContainerCall{
				Files: map[string]interface{}{
					// implicitly bound
					fileName: "",
				},
			}

			provideddataDirPath := "dummydataDirPath"
			providedContainerID := "dummyContainerID"
			providedRootOpID := "dummyRootOpID"
			providedOpHandle := new(data.FakeHandle)

			expectedScratchDirPath := filepath.Join(
				provideddataDirPath,
				"dcg",
				providedRootOpID,
				"containers",
				providedContainerID,
				"fs",
			)

			fakeFilesInterpreter := new(files.FakeInterpreter)

			objectUnderTest := _interpreter{
				dirsInterpreter:    new(dirs.FakeInterpreter),
				envVarsInterpreter: new(envvars.FakeInterpreter),
				filesInterpreter:   fakeFilesInterpreter,
				imageInterpreter:   new(image.FakeInterpreter),
				os:                 new(ios.Fake),
				dataDirPath:        provideddataDirPath,
				socketsInterpreter: new(sockets.FakeInterpreter),
			}

			/* act */
			objectUnderTest.Interpret(
				providedScope,
				providedSCGContainerCall,
				providedContainerID,
				providedRootOpID,
				providedOpHandle,
			)

			/* assert */
			actualOpHandle, actualScope, actualScgContainerCallFiles, actualScratchDir := fakeFilesInterpreter.InterpretArgsForCall(0)
			Expect(actualOpHandle).To(Equal(providedOpHandle))
			Expect(actualScope).To(Equal(providedScope))
			Expect(actualScgContainerCallFiles).To(Equal(providedSCGContainerCall.Files))
			Expect(actualScratchDir).To(Equal(expectedScratchDirPath))
		})
		Context("files.Interpret errors", func() {
			It("should return expected error", func() {
				/* arrange */
				expectedErr := errors.New("dummyError")
				fakeFilesInterpreter := new(files.FakeInterpreter)
				fakeFilesInterpreter.InterpretReturns(nil, expectedErr)

				objectUnderTest := _interpreter{
					dirsInterpreter:    new(dirs.FakeInterpreter),
					envVarsInterpreter: new(envvars.FakeInterpreter),
					filesInterpreter:   fakeFilesInterpreter,
					imageInterpreter:   new(image.FakeInterpreter),
					os:                 new(ios.Fake),
				}

				/* act */
				_, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGContainerCall{},
					"dummyContainerID",
					"dummyRootOpID",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
		Context("files.Interpret doesn't error", func() {
			It("should return expected dcgContainerCall.Files", func() {
				/* arrange */
				expectedDCGContainerCallFiles := map[string]string{
					"dummyName": "dummyValue",
				}

				fakeFilesInterpreter := new(files.FakeInterpreter)
				fakeFilesInterpreter.InterpretReturns(expectedDCGContainerCallFiles, nil)

				objectUnderTest := _interpreter{
					dirsInterpreter:    new(dirs.FakeInterpreter),
					envVarsInterpreter: new(envvars.FakeInterpreter),
					filesInterpreter:   fakeFilesInterpreter,
					imageInterpreter:   new(image.FakeInterpreter),
					os:                 new(ios.Fake),
					socketsInterpreter: new(sockets.FakeInterpreter),
				}

				/* act */
				actualResult, _ := objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGContainerCall{},
					"dummyContainerID",
					"dummyRootOpID",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualResult.Files).To(Equal(expectedDCGContainerCallFiles))
			})
		})

		It("should call image w/ expected args", func() {
			/* arrange */
			containerFileBind := "dummyContainerFileBind"

			providedScope := map[string]*model.Value{
				containerFileBind: {String: new(string)},
			}
			providedSCGContainerCall := &model.SCGContainerCall{
				Image: &model.SCGContainerCallImage{
					Ref: "dummyImageRef",
					PullCreds: &model.SCGPullCreds{
						Username: "dummyUsername",
						Password: "dummyPassword",
					},
				},
			}

			providedOpHandle := new(data.FakeHandle)

			fakeImageInterpreter := new(image.FakeInterpreter)

			objectUnderTest := _interpreter{
				dirsInterpreter:    new(dirs.FakeInterpreter),
				envVarsInterpreter: new(envvars.FakeInterpreter),
				filesInterpreter:   new(files.FakeInterpreter),
				imageInterpreter:   fakeImageInterpreter,
				os:                 new(ios.Fake),
				socketsInterpreter: new(sockets.FakeInterpreter),
			}

			/* act */
			objectUnderTest.Interpret(
				providedScope,
				providedSCGContainerCall,
				"dummyContainerID",
				"dummyRootOpID",
				providedOpHandle,
			)

			/* assert */
			actualScope,
				actualScgContainerCallImage,
				actualOpHandle := fakeImageInterpreter.InterpretArgsForCall(0)

			Expect(actualScope).To(Equal(providedScope))
			Expect(actualScgContainerCallImage).To(Equal(providedSCGContainerCall.Image))
			Expect(actualOpHandle).To(Equal(providedOpHandle))

		})
		Context("image.Interpret errors", func() {
			It("should return expected error", func() {
				/* arrange */
				expectedErr := errors.New("dummyError")
				fakeImageInterpreter := new(image.FakeInterpreter)
				fakeImageInterpreter.InterpretReturns(nil, expectedErr)

				objectUnderTest := _interpreter{
					dirsInterpreter:    new(dirs.FakeInterpreter),
					envVarsInterpreter: new(envvars.FakeInterpreter),
					filesInterpreter:   new(files.FakeInterpreter),
					imageInterpreter:   fakeImageInterpreter,
					os:                 new(ios.Fake),
				}

				/* act */
				_, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGContainerCall{},
					"dummyContainerID",
					"dummyRootOpID",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
		Context("image.Interpret doesn't error", func() {
			It("should return expected dcgContainerCall.Image", func() {
				/* arrange */
				expectedDCGContainerCallImage := &model.DCGContainerCallImage{
					Ref: "dummyImageRef",
					PullCreds: &model.PullCreds{
						Username: "dummyUsername",
						Password: "dummyPassword",
					},
				}

				fakeImageInterpreter := new(image.FakeInterpreter)
				fakeImageInterpreter.InterpretReturns(expectedDCGContainerCallImage, nil)

				objectUnderTest := _interpreter{
					dirsInterpreter:    new(dirs.FakeInterpreter),
					envVarsInterpreter: new(envvars.FakeInterpreter),
					filesInterpreter:   new(files.FakeInterpreter),
					imageInterpreter:   fakeImageInterpreter,
					os:                 new(ios.Fake),
					socketsInterpreter: new(sockets.FakeInterpreter),
				}

				/* act */
				actualResult, _ := objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGContainerCall{},
					"dummyContainerID",
					"dummyRootOpID",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualResult.Image).To(Equal(expectedDCGContainerCallImage))
			})
		})

		It("should call sockets w/ expected args", func() {
			/* arrange */
			containerFileBind := "dummyContainerFileBind"

			providedScope := map[string]*model.Value{
				containerFileBind: {String: new(string)},
			}

			envVarName := "dummyEnvVarName"
			providedSCGContainerCall := &model.SCGContainerCall{
				Sockets: map[string]string{
					// implicitly bound
					envVarName: "",
				},
			}

			provideddataDirPath := "dummydataDirPath"
			providedContainerID := "dummyContainerID"
			providedRootOpID := "dummyRootOpID"

			expectedScratchDirPath := filepath.Join(
				provideddataDirPath,
				"dcg",
				providedRootOpID,
				"containers",
				providedContainerID,
				"fs",
			)

			fakeSocketsInterpreter := new(sockets.FakeInterpreter)

			objectUnderTest := _interpreter{
				dirsInterpreter:    new(dirs.FakeInterpreter),
				envVarsInterpreter: new(envvars.FakeInterpreter),
				filesInterpreter:   new(files.FakeInterpreter),
				imageInterpreter:   new(image.FakeInterpreter),
				os:                 new(ios.Fake),
				dataDirPath:        provideddataDirPath,
				socketsInterpreter: fakeSocketsInterpreter,
			}

			/* act */
			objectUnderTest.Interpret(
				providedScope,
				providedSCGContainerCall,
				providedContainerID,
				providedRootOpID,
				new(data.FakeHandle),
			)

			/* assert */
			actualScope, actualScgContainerCallSockets, actualScratchDir := fakeSocketsInterpreter.InterpretArgsForCall(0)
			Expect(actualScope).To(Equal(providedScope))
			Expect(actualScgContainerCallSockets).To(Equal(providedSCGContainerCall.Sockets))
			Expect(actualScratchDir).To(Equal(expectedScratchDirPath))
		})
		Context("sockets.Interpret errors", func() {
			It("should return expected error", func() {
				/* arrange */
				expectedErr := errors.New("dummyError")
				fakeSocketsInterpreter := new(sockets.FakeInterpreter)
				fakeSocketsInterpreter.InterpretReturns(nil, expectedErr)

				objectUnderTest := _interpreter{
					dirsInterpreter:    new(dirs.FakeInterpreter),
					envVarsInterpreter: new(envvars.FakeInterpreter),
					filesInterpreter:   new(files.FakeInterpreter),
					imageInterpreter:   new(image.FakeInterpreter),
					os:                 new(ios.Fake),
					socketsInterpreter: fakeSocketsInterpreter,
				}

				/* act */
				_, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGContainerCall{},
					"dummyContainerID",
					"dummyRootOpID",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
		Context("sockets.Interpret doesn't error", func() {
			It("should return expected dcgContainerCall.Sockets", func() {
				/* arrange */
				expectedDCGContainerCallSockets := map[string]string{
					"dummyName": "dummyValue",
				}

				fakeSocketsInterpreter := new(sockets.FakeInterpreter)
				fakeSocketsInterpreter.InterpretReturns(expectedDCGContainerCallSockets, nil)

				objectUnderTest := _interpreter{
					dirsInterpreter:    new(dirs.FakeInterpreter),
					envVarsInterpreter: new(envvars.FakeInterpreter),
					filesInterpreter:   new(files.FakeInterpreter),
					imageInterpreter:   new(image.FakeInterpreter),
					os:                 new(ios.Fake),
					socketsInterpreter: fakeSocketsInterpreter,
				}

				/* act */
				actualResult, _ := objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGContainerCall{},
					"dummyContainerID",
					"dummyRootOpID",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualResult.Sockets).To(Equal(expectedDCGContainerCallSockets))
			})
		})
	})
})
