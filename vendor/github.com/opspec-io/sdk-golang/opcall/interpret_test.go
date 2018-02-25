package opcall

import (
	"errors"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/expression"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/opcall/inputs"
	"github.com/opspec-io/sdk-golang/pkg"
	"github.com/opspec-io/sdk-golang/util/uniquestring"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var _ = Context("OpCall", func() {
	Context("Interpret", func() {
		Context("called w/ opspec test-suite scenarios", func() {
			It("should return result fulfilling scenario.call.expect", func() {
				tempDir, err := ioutil.TempDir("", "")
				if nil != err {
					panic(err)
				}
				rootPath := "../github.com/opspec-io/test-suite/scenarios/pkg"

				filepath.Walk(rootPath,
					func(path string, info os.FileInfo, err error) error {
						if info.IsDir() {
							scenariosDotYmlFilePath := filepath.Join(path, "scenarios.yml")
							if _, err := os.Stat(scenariosDotYmlFilePath); nil == err {
								/* arrange */
								absPkgPath, err := filepath.Abs(path)
								if nil != err {
									panic(fmt.Errorf("error getting absPkgPath %v; error was %v", path, err))
								}

								pkg := pkg.New()
								pkgHandle, err := pkg.Resolve(
									absPkgPath,
									pkg.NewFSProvider(),
								)
								if nil != err {
									panic(fmt.Errorf("error resolving pkg for %v; error was %v", path, err))
								}

								scenariosDotYmlBytes, err := ioutil.ReadFile(scenariosDotYmlFilePath)
								if nil != err {
									panic(err)
								}

								var scenarioDotYml []struct {
									Name      string
									Interpret *struct {
										Expect string
										Scope  map[string]*model.Value
									}
								}
								if err := yaml.Unmarshal(scenariosDotYmlBytes, &scenarioDotYml); nil != err {
									panic(fmt.Errorf("error unmarshalling scenario.yml for %v; error was %v", path, err))
								}

								for _, scenario := range scenarioDotYml {
									if nil != scenario.Interpret {
										scgOpCall := &model.SCGOpCall{
											Pkg: &model.SCGOpCallPkg{
												Ref: absPkgPath,
											},
											Inputs: map[string]interface{}{},
										}

										for name := range scenario.Interpret.Scope {
											// map as passed
											scgOpCall.Inputs[name] = ""
										}

										/* act */
										objectUnderTest := New(tempDir)
										_, actualErr := objectUnderTest.Interpret(
											scenario.Interpret.Scope,
											scgOpCall,
											"",
											pkgHandle,
											"",
										)

										/* assert */
										description := fmt.Sprintf("pkg: '%v'\nscenario: '%v'", path, scenario.Name)
										switch expect := scenario.Interpret.Expect; expect {
										case "success":
											Expect(actualErr).To(BeNil(), description)
										case "failure":
											Expect(actualErr).To(Not(BeNil()), description)
										}
									}
								}
							}
						}
						return nil
					})
			})
		})
		It("should call pkg.NewFSProvider w/ expected args", func() {
			/* arrange */
			providedParentPkgHandle := new(pkg.FakeHandle)
			providedParentPkgHandle.PathReturns(new(string))

			fakePkg := new(pkg.Fake)
			// error to trigger immediate return
			fakePkg.ResolveReturns(nil, errors.New("dummyError"))

			objectUnderTest := _OpCall{
				expression: new(expression.Fake),
				pkg:        fakePkg,
			}

			/* act */
			objectUnderTest.Interpret(
				map[string]*model.Value{},
				&model.SCGOpCall{
					Pkg: &model.SCGOpCallPkg{
						Ref: "dummyPkgRef",
					},
				},
				"dummyOpId",
				providedParentPkgHandle,
				"dummyRootOpId",
			)

			/* assert */
			Expect(fakePkg.NewFSProviderArgsForCall(0)).To(ConsistOf(filepath.Dir(providedParentPkgHandle.Ref())))
		})
		Context("scgOpCall.Pkg.PullCreds is nil", func() {
			It("should call pkg.NewGitProvider w/ expected args", func() {
				/* arrange */
				providedParentPkgHandle := new(pkg.FakeHandle)
				providedParentPkgHandle.PathReturns(new(string))

				providedPkgCachePath := "dummyPkgCachePath"

				fakePkg := new(pkg.Fake)
				// error to trigger immediate return
				fakePkg.ResolveReturns(nil, errors.New("dummyError"))

				objectUnderTest := _OpCall{
					pkg:          fakePkg,
					pkgCachePath: providedPkgCachePath,
				}

				/* act */
				objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGOpCall{
						Pkg: &model.SCGOpCallPkg{
							Ref: "dummyPkgRef",
						},
					},
					"dummyOpId",
					providedParentPkgHandle,
					"dummyRootOpId",
				)

				/* assert */
				actualBasePath,
					actualPullCreds := fakePkg.NewGitProviderArgsForCall(0)

				Expect(actualBasePath).To(Equal(providedPkgCachePath))
				Expect(actualPullCreds).To(BeNil())
			})
		})
		Context("scgOpCall.Pkg.PullCreds isn't nil", func() {
			Context("string.Interpret errs", func() {
				It("should return expected result", func() {
					/* arrange */
					fakeExpression := new(expression.Fake)
					interpretError := errors.New("dummyError")
					fakeExpression.EvalToStringReturns(nil, interpretError)

					objectUnderTest := _OpCall{
						expression: fakeExpression,
					}

					/* act */
					_, actualError := objectUnderTest.Interpret(
						map[string]*model.Value{},
						&model.SCGOpCall{
							Pkg: &model.SCGOpCallPkg{
								PullCreds: &model.SCGPullCreds{},
							},
						},
						"dummyOpId",
						new(pkg.FakeHandle),
						"dummyRootOpId",
					)

					/* assert */
					Expect(actualError).To(Equal(interpretError))
				})
			})
			Context("string.Interpret doesn't err", func() {
				It("should call pkg.NewGitProvider w/ expected args", func() {
					/* arrange */
					providedParentPkgHandle := new(pkg.FakeHandle)
					providedParentPkgHandle.PathReturns(new(string))

					providedPkgCachePath := "dummyPkgCachePath"

					fakeExpression := new(expression.Fake)
					expectedPullCreds := &model.PullCreds{Username: "dummyUsername", Password: "dummyPassword"}
					fakeExpression.EvalToStringReturnsOnCall(0, &model.Value{String: &expectedPullCreds.Username}, nil)
					fakeExpression.EvalToStringReturnsOnCall(1, &model.Value{String: &expectedPullCreds.Password}, nil)

					fakePkg := new(pkg.Fake)
					// error to trigger immediate return
					fakePkg.ResolveReturns(nil, errors.New("dummyError"))

					objectUnderTest := _OpCall{
						expression:   fakeExpression,
						pkg:          fakePkg,
						pkgCachePath: providedPkgCachePath,
					}

					/* act */
					objectUnderTest.Interpret(
						map[string]*model.Value{},
						&model.SCGOpCall{
							Pkg: &model.SCGOpCallPkg{
								Ref:       "dummyPkgRef",
								PullCreds: &model.SCGPullCreds{},
							},
						},
						"dummyOpId",
						providedParentPkgHandle,
						"dummyRootOpId",
					)

					/* assert */
					actualBasePath,
						actualPullCreds := fakePkg.NewGitProviderArgsForCall(0)

					Expect(actualBasePath).To(Equal(providedPkgCachePath))
					Expect(actualPullCreds).To(Equal(expectedPullCreds))
				})
			})
		})
		It("should call pkg.Resolve w/ expected args", func() {
			/* arrange */
			providedParentPkgHandle := new(pkg.FakeHandle)
			providedParentPkgHandle.PathReturns(new(string))

			providedRootFSPath := "dummyRootFSPath"
			providedSCGOpCall := &model.SCGOpCall{
				Pkg: &model.SCGOpCallPkg{
					Ref: "dummyPkgRef",
				},
			}

			expectedPkgRef := providedSCGOpCall.Pkg.Ref

			fakePkg := new(pkg.Fake)

			expectedPkgProviders := []pkg.Provider{
				new(pkg.FakeProvider),
				new(pkg.FakeProvider),
			}
			fakePkg.NewFSProviderReturns(expectedPkgProviders[0])
			fakePkg.NewGitProviderReturns(expectedPkgProviders[1])

			// error to trigger immediate return
			fakePkg.ResolveReturns(nil, errors.New("dummyError"))

			objectUnderTest := _OpCall{
				pkg:          fakePkg,
				pkgCachePath: filepath.Join(providedRootFSPath, "pkgs"),
			}

			/* act */
			objectUnderTest.Interpret(
				map[string]*model.Value{},
				providedSCGOpCall,
				"dummyOpId",
				providedParentPkgHandle,
				"dummyRootOpId",
			)

			/* assert */
			actualPkgRef, actualPkgProviders := fakePkg.ResolveArgsForCall(0)
			Expect(actualPkgRef).To(Equal(expectedPkgRef))
			Expect(actualPkgProviders).To(Equal(expectedPkgProviders))
		})
		Context("pkg.Resolve errs", func() {
			It("should return err", func() {
				/* arrange */
				providedParentPkgHandle := new(pkg.FakeHandle)
				providedParentPkgHandle.PathReturns(new(string))

				expectedErr := errors.New("dummyError")
				fakePkg := new(pkg.Fake)
				fakePkg.ResolveReturns(nil, expectedErr)

				objectUnderTest := _OpCall{
					pkg:                 fakePkg,
					uniqueStringFactory: new(uniquestring.Fake),
				}

				/* act */
				_, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGOpCall{Pkg: &model.SCGOpCallPkg{}},
					"dummyOpId",
					providedParentPkgHandle,
					"dummyRootOpId",
				)

				/* assert */
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
		Context("pkg.Resolve doesn't err", func() {
			It("should call pkg.GetManifest w/ expected args", func() {
				/* arrange */
				providedParentPkgHandle := new(pkg.FakeHandle)
				providedParentPkgHandle.PathReturns(new(string))

				fakePkgHandle := new(pkg.FakeHandle)

				fakePkg := new(pkg.Fake)
				fakePkg.ResolveReturns(fakePkgHandle, nil)

				expectedErr := errors.New("dummyError")
				// err to trigger immediate return
				fakePkg.GetManifestReturns(nil, expectedErr)

				objectUnderTest := _OpCall{
					pkg:                 fakePkg,
					uniqueStringFactory: new(uniquestring.Fake),
				}

				/* act */
				objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGOpCall{Pkg: &model.SCGOpCallPkg{}},
					"dummyOpId",
					providedParentPkgHandle,
					"dummyRootOpId",
				)

				/* assert */
				actualHandle := fakePkg.GetManifestArgsForCall(0)
				Expect(actualHandle).To(Equal(fakePkgHandle))
			})
			Context("pkg.GetManifest errs", func() {
				It("should return err", func() {
					/* arrange */
					providedParentPkgHandle := new(pkg.FakeHandle)
					providedParentPkgHandle.PathReturns(new(string))

					expectedErr := errors.New("dummyError")
					fakePkg := new(pkg.Fake)
					fakePkg.GetManifestReturns(nil, expectedErr)

					objectUnderTest := _OpCall{
						pkg:                 fakePkg,
						uniqueStringFactory: new(uniquestring.Fake),
					}

					/* act */
					_, actualErr := objectUnderTest.Interpret(
						map[string]*model.Value{},
						&model.SCGOpCall{Pkg: &model.SCGOpCallPkg{}},
						"dummyOpId",
						providedParentPkgHandle,
						"dummyRootOpId",
					)

					/* assert */
					Expect(actualErr).To(Equal(expectedErr))
				})
			})
			Context("pkg.GetManifest doesn't err", func() {
				It("should call inputs.Interpret w/ expected inputs", func() {
					/* arrange */
					providedScope := map[string]*model.Value{
						"dummyScopeRef1Name": {String: new(string)},
					}
					expectedScope := providedScope

					expectedInputArgs := map[string]interface{}{"dummySCGOpCallInputName": "dummyScgOpCallInputValue"}

					providedSCGOpCall := &model.SCGOpCall{
						Inputs: expectedInputArgs,
						Pkg:    &model.SCGOpCallPkg{},
					}

					providedOpId := "dummyOpId"

					providedParentPkgHandle := new(pkg.FakeHandle)
					parentPkgPath := "dummyParentPkgPath"
					providedParentPkgHandle.PathReturns(&parentPkgPath)

					fakePkgHandle := new(pkg.FakeHandle)
					pkgPath := "dummyPkgPath"
					fakePkgHandle.PathReturns(&pkgPath)

					fakePkg := new(pkg.Fake)
					fakePkg.ResolveReturns(fakePkgHandle, nil)

					expectedInputParams := map[string]*model.Param{
						"dummyParam1Name": {String: &model.StringParam{}},
					}

					returnedManifest := &model.PkgManifest{
						Inputs: expectedInputParams,
					}
					fakePkg.GetManifestReturns(returnedManifest, nil)

					fakeInputs := new(inputs.Fake)

					dcgScratchDir := "dummyDCGScratchDir"

					objectUnderTest := _OpCall{
						dcgScratchDir:       dcgScratchDir,
						pkg:                 fakePkg,
						uniqueStringFactory: new(uniquestring.Fake),
						inputs:              fakeInputs,
					}

					/* act */
					objectUnderTest.Interpret(
						providedScope,
						providedSCGOpCall,
						providedOpId,
						providedParentPkgHandle,
						"dummyRootOpId",
					)

					/* assert */
					actualSCGArgs,
						actualSCGInputs,
						actualParentPkgHandle,
						actualPkgRef,
						actualScope,
						actualOpScratchDir := fakeInputs.InterpretArgsForCall(0)

					Expect(actualScope).To(Equal(expectedScope))
					Expect(actualSCGArgs).To(Equal(expectedInputArgs))
					Expect(actualParentPkgHandle).To(Equal(providedParentPkgHandle))
					Expect(actualPkgRef).To(Equal(pkgPath))
					Expect(actualSCGInputs).To(Equal(expectedInputParams))
					Expect(actualOpScratchDir).To(Equal(filepath.Join(dcgScratchDir, providedOpId)))

				})
			})
		})
	})
})
