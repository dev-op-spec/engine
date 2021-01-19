package clioutput

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/cli/internal/clicolorer"
	clicolorerFakes "github.com/opctl/opctl/cli/internal/clicolorer/fakes"
	"github.com/opctl/opctl/sdks/go/model"
)

var _ = Context("output", func() {
	var opFormatter SimpleOpFormatter
	Context("newOutput", func() {
		It("should return output", func() {
			/* arrange/act/assert */
			Expect(New(
				new(clicolorerFakes.FakeCliColorer),
				opFormatter,
				new(fakeWriter),
				new(fakeWriter),
			)).To(Not(BeNil()))
		})
	})
	_cliColorer := clicolorer.New()
	Context("Attention", func() {
		providedFormat := "dummyFormat %v %v"
		It("should call stdWriter w/ expected args", func() {
			/* arrange */
			expectedWriteArg := fmt.Sprintln(_cliColorer.Attention(providedFormat))

			fakeStdWriter := new(fakeWriter)
			objectUnderTest, err := New(
				_cliColorer,
				opFormatter,
				new(fakeWriter),
				fakeStdWriter,
			)
			Expect(err).To(BeNil())

			/* act */
			objectUnderTest.Attention(providedFormat)

			/* assert */
			Expect(string(fakeStdWriter.WriteArgsForCall(0))).
				To(Equal(expectedWriteArg))
		})
	})
	Context("Warning", func() {
		providedFormat := "dummyFormat %v %v"
		It("should call stdWriter w/ expected args", func() {
			/* arrange */
			expectedWriteArg := fmt.Sprintln(_cliColorer.Error(providedFormat))

			fakeErrWriter := new(fakeWriter)
			objectUnderTest, err := New(
				_cliColorer,
				opFormatter,
				new(fakeWriter),
				fakeErrWriter,
			)
			Expect(err).To(BeNil())

			/* act */
			objectUnderTest.Warning(providedFormat)

			/* assert */
			Expect(string(fakeErrWriter.WriteArgsForCall(0))).
				To(Equal(expectedWriteArg))
		})
	})
	Context("Error", func() {
		providedFormat := "dummyFormat %v %v"
		It("should call errWriter w/ expected args", func() {
			/* arrange */
			expectedWriteArg := fmt.Sprintln(_cliColorer.Error(providedFormat))

			fakeErrWriter := new(fakeWriter)
			objectUnderTest, err := New(
				_cliColorer,
				opFormatter,
				fakeErrWriter,
				new(fakeWriter),
			)
			Expect(err).To(BeNil())

			/* act */
			objectUnderTest.Error(providedFormat)

			/* assert */
			Expect(string(fakeErrWriter.WriteArgsForCall(0))).
				To(Equal(expectedWriteArg))
		})
	})
	Context("Event", func() {
		Context("ContainerStdErrWrittenTo", func() {
			It("should call stdWriter w/ expected args", func() {
				/* arrange */
				providedEvent := &model.Event{
					ContainerStdErrWrittenTo: &model.ContainerStdErrWrittenTo{
						ContainerID: "acontainerid",
						Data:        []byte("dummyData"),
					},
					Timestamp: time.Now(),
				}
				expectedWriteArg := "\x1b[2m[acontain]\x1b[0m " + string(providedEvent.ContainerStdErrWrittenTo.Data)

				fakeErrWriter := new(fakeWriter)
				objectUnderTest, err := New(
					_cliColorer,
					opFormatter,
					fakeErrWriter,
					new(fakeWriter),
				)
				Expect(err).To(BeNil())

				/* act */
				objectUnderTest.Event(providedEvent)

				/* assert */
				Expect(string(fakeErrWriter.WriteArgsForCall(0))).
					To(Equal(expectedWriteArg))
			})
		})
		Context("ContainerOutWrittenTo", func() {
			It("should call stdWriter w/ expected args", func() {
				/* arrange */
				providedEvent := &model.Event{
					ContainerStdOutWrittenTo: &model.ContainerStdOutWrittenTo{
						ContainerID: "acontainerid",
						Data:        []byte("dummyData"),
					},
					Timestamp: time.Now(),
				}
				expectedWriteArg := "\x1b[2m[acontain]\x1b[0m " + string(providedEvent.ContainerStdOutWrittenTo.Data)

				fakeStdWriter := new(fakeWriter)
				objectUnderTest, err := New(
					_cliColorer,
					opFormatter,
					new(fakeWriter),
					fakeStdWriter,
				)
				Expect(err).To(BeNil())

				/* act */
				objectUnderTest.Event(providedEvent)

				/* assert */
				Expect(string(fakeStdWriter.WriteArgsForCall(0))).
					To(Equal(expectedWriteArg))
			})
		})
		Context("CallEnded", func() {
			Context("Call.Container truthy", func() {
				It("should call stdWriter w/ expected args", func() {
					/* arrange */
					imageRef := "imageRef"
					providedEvent := &model.Event{
						CallEnded: &model.CallEnded{
							Call: model.Call{
								Container: &model.ContainerCall{
									Image: &model.ContainerCallImage{
										Ref: &imageRef,
									},
								},
								ID: "acontainerID",
							},
							Outcome: model.OpOutcomeSucceeded,
							Ref:     "ref",
						},
						Timestamp: time.Now(),
					}
					expectedWriteArg := "\x1b[2m[acontain ref]\x1b[0m \x1b[92;1mexited\x1b[0m imageRef\n"

					fakeStdWriter := new(fakeWriter)
					objectUnderTest, err := New(
						_cliColorer,
						opFormatter,
						new(fakeWriter),
						fakeStdWriter,
					)
					Expect(err).To(BeNil())

					/* act */
					objectUnderTest.Event(providedEvent)

					/* assert */
					Expect(string(fakeStdWriter.WriteArgsForCall(0))).
						To(Equal(expectedWriteArg))
				})
			})
			Context("Call.Op truthy", func() {

				Context("Outcome==FAILED", func() {
					It("should call errWriter w/ expected args", func() {
						/* arrange */
						providedEvent := &model.Event{
							CallEnded: &model.CallEnded{
								Call: model.Call{
									ID: "thisisacallID",
									Op: &model.OpCall{
										BaseCall: model.BaseCall{
											OpPath: "opPath",
										},
									},
								},
								Error: &model.CallEndedError{
									Message: "message",
								},
								Ref:     "ref",
								Outcome: "FAILED",
							},
							Timestamp: time.Now(),
						}
						expectedWriteArg := "[2m[thisisac opPath][0m [91;1mended with error[0m\nmessage\n"

						fakeErrWriter := new(fakeWriter)
						objectUnderTest, err := New(
							_cliColorer,
							opFormatter,
							fakeErrWriter,
							new(fakeWriter),
						)
						Expect(err).To(BeNil())

						/* act */
						objectUnderTest.Event(providedEvent)

						/* assert */
						Expect(string(fakeErrWriter.WriteArgsForCall(0))).
							To(Equal(expectedWriteArg))
					})
				})
				Context("Outcome==SUCCEEDED", func() {
					It("should call stdWriter w/ expected args", func() {
						/* arrange */
						providedEvent := &model.Event{
							CallEnded: &model.CallEnded{
								Call: model.Call{
									ID: "thisisacallID",
									Op: &model.OpCall{
										BaseCall: model.BaseCall{
											OpPath: "opPath",
										},
									},
								},
								Ref:     "ref",
								Outcome: "SUCCEEDED",
							},
							Timestamp: time.Now(),
						}
						expectedWriteArg := "\x1b[2m[thisisac opPath]\x1b[0m \x1b[92;1mended\x1b[0m\n"

						fakeStdWriter := new(fakeWriter)
						objectUnderTest, err := New(
							_cliColorer,
							opFormatter,
							new(fakeWriter),
							fakeStdWriter,
						)
						Expect(err).To(BeNil())

						/* act */
						objectUnderTest.Event(providedEvent)

						/* assert */
						Expect(string(fakeStdWriter.WriteArgsForCall(0))).
							To(Equal(expectedWriteArg))
					})
				})
				Context("Outcome==KILLED", func() {
					It("should call stdWriter w/ expected args", func() {
						/* arrange */
						providedEvent := &model.Event{
							CallEnded: &model.CallEnded{
								Call: model.Call{
									ID: "thisisacallID",
									Op: &model.OpCall{
										BaseCall: model.BaseCall{
											OpPath: "opPath",
										},
									},
								},
								Ref:     "ref",
								Outcome: "KILLED",
							},
							Timestamp: time.Now(),
						}
						expectedWriteArg := "\x1b[2m[thisisac opPath]\x1b[0m \x1b[96;1mended\x1b[0m\n"

						fakeStdWriter := new(fakeWriter)
						objectUnderTest, err := New(
							_cliColorer,
							opFormatter,
							new(fakeWriter),
							fakeStdWriter,
						)
						Expect(err).To(BeNil())

						/* act */
						objectUnderTest.Event(providedEvent)

						/* assert */
						Expect(string(fakeStdWriter.WriteArgsForCall(0))).
							To(Equal(expectedWriteArg))
					})
				})
			})
			Context("Error truthy", func() {
				It("should call errWriter w/ expected args", func() {
					/* arrange */
					providedEvent := &model.Event{
						CallEnded: &model.CallEnded{
							Call: model.Call{
								ID: "thisisacallID",
							},
							Error: &model.CallEndedError{
								Message: "message",
							},
							Ref: "ref",
						},
						Timestamp: time.Now(),
					}
					expectedWriteArg := "\x1b[2m[thisisac ref]\x1b[0m \x1b[91;1mmessage\x1b[0m\n"

					fakeErrWriter := new(fakeWriter)
					objectUnderTest, err := New(
						_cliColorer,
						opFormatter,
						fakeErrWriter,
						new(fakeWriter),
					)
					Expect(err).To(BeNil())

					/* act */
					objectUnderTest.Event(providedEvent)

					/* assert */
					Expect(string(fakeErrWriter.WriteArgsForCall(0))).
						To(Equal(expectedWriteArg))
				})
			})
		})
		Context("CallStarted", func() {
			Context("Call.Container truthy", func() {
				It("should call stdWriter w/ expected args", func() {
					/* arrange */
					imageRef := "imageRef"
					providedEvent := &model.Event{
						CallStarted: &model.CallStarted{
							Call: model.Call{
								ID: "thisisacallID",
								Container: &model.ContainerCall{
									Image: &model.ContainerCallImage{
										Ref: &imageRef,
									},
								},
							},
							Ref: "ref",
						},
						Timestamp: time.Now(),
					}
					expectedWriteArg := "[2m[thisisac ref][0m [96;1mstarted container[0m imageRef\n"

					fakeStdWriter := new(fakeWriter)
					objectUnderTest, err := New(
						_cliColorer,
						opFormatter,
						new(fakeWriter),
						fakeStdWriter,
					)
					Expect(err).To(BeNil())

					/* act */
					objectUnderTest.Event(providedEvent)

					/* assert */
					Expect(string(fakeStdWriter.WriteArgsForCall(0))).
						To(Equal(expectedWriteArg))
				})
			})
			Describe("Call.Op truthy", func() {
				It("should call stdWriter w/ expected args", func() {
					/* arrange */
					providedEvent := &model.Event{
						CallStarted: &model.CallStarted{
							Call: model.Call{
								ID: "thisisacallID",
								Op: &model.OpCall{
									BaseCall: model.BaseCall{
										OpPath: "opPath",
									},
								},
							},
							Ref: "ref",
						},
						Timestamp: time.Now(),
					}
					expectedWriteArg := "\x1b[2m[thisisac opPath]\x1b[0m \x1b[96;1mstarted op\x1b[0m\n"

					fakeStdWriter := new(fakeWriter)
					objectUnderTest, err := New(
						_cliColorer,
						opFormatter,
						new(fakeWriter),
						fakeStdWriter,
					)
					Expect(err).To(BeNil())

					/* act */
					objectUnderTest.Event(providedEvent)

					/* assert */
					Expect(string(fakeStdWriter.WriteArgsForCall(0))).
						To(Equal(expectedWriteArg))
				})
			})
		})
	})
	Context("Success", func() {
		providedFormat := "dummyFormat %v %v"
		It("should call stdWriter w/ expected args", func() {
			/* arrange */
			expectedWriteArg := fmt.Sprintln(_cliColorer.Success(providedFormat))

			fakeStdWriter := new(fakeWriter)
			objectUnderTest, err := New(
				_cliColorer,
				opFormatter,
				new(fakeWriter),
				fakeStdWriter,
			)
			Expect(err).To(BeNil())

			/* act */
			objectUnderTest.Success(providedFormat)

			/* assert */
			Expect(string(fakeStdWriter.WriteArgsForCall(0))).
				To(Equal(expectedWriteArg))
		})
	})
	Context("DisableColor", func() {
		It("should call colorer to disable color", func() {
			/* arrange */
			fakeCliColorer := new(clicolorerFakes.FakeCliColorer)

			objectUnderTest, err := New(
				fakeCliColorer,
				opFormatter,
				new(fakeWriter),
				new(fakeWriter),
			)
			Expect(err).To(BeNil())

			/* act */
			objectUnderTest.DisableColor()

			/* assert */
			Expect(fakeCliColorer.DisableColorCallCount()).To(Equal(1))
		})
	})
})
