package core

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/sdk-golang/data"
	"github.com/opctl/sdk-golang/model"
	"github.com/opctl/sdk-golang/opspec/interpreter/call"
	"github.com/opctl/sdk-golang/util/pubsub"
)

var _ = Context("caller", func() {
	Context("newCaller", func() {
		It("should return caller", func() {
			/* arrange/act/assert */
			Expect(
				newCaller(
					new(call.FakeInterpreter),
					new(fakeContainerCaller),
					"dummyDataDir",
					new(fakeCallStore),
					new(fakeCallKiller),
					new(pubsub.Fake),
				),
			).To(Not(BeNil()))
		})
	})
	Context("Call", func() {
		closedEventChan := make(chan model.Event, 1000)
		close(closedEventChan)

		Context("Null SCG", func() {
			It("should not throw", func() {
				/* arrange */
				fakeContainerCaller := new(fakeContainerCaller)

				/* act */
				objectUnderTest := _caller{
					callInterpreter: new(call.FakeInterpreter),
					callStore:       new(fakeCallStore),
					containerCaller: fakeContainerCaller,
					pubSub:          new(pubsub.Fake),
				}

				/* assert */
				objectUnderTest.Call(
					context.Background(),
					"dummyCallID",
					map[string]*model.Value{},
					nil,
					new(data.FakeHandle),
					nil,
					"dummyRootOpID",
				)
			})
		})

		Context("nil != scg.Loop", func() {
			It("should call looper.Loop w/ expected args", func() {
				/* arrange */
				providedCallID := "providedCallID"
				providedScope := map[string]*model.Value{}
				providedSCG := &model.SCG{
					Loop: &model.SCGLoop{},
				}
				providedOpHandle := new(data.FakeHandle)
				providedParentIDValue := "providedParentID"
				providedParentID := &providedParentIDValue
				providedRootOpID := "providedRootOpID"

				fakePubSub := new(pubsub.Fake)
				// ensure eventChan closed so call exits
				fakePubSub.SubscribeReturns(closedEventChan, nil)

				fakeLooper := new(fakeLooper)
				objectUnderTest := _caller{
					callStore: new(fakeCallStore),
					looper:    fakeLooper,
					pubSub:    fakePubSub,
				}

				/* act */
				objectUnderTest.Call(
					context.Background(),
					providedCallID,
					providedScope,
					providedSCG,
					providedOpHandle,
					providedParentID,
					providedRootOpID,
				)

				/* assert */
				_,
					actualID,
					actualScope,
					actualSCG,
					actualOpHandle,
					actualParentID,
					actualRootOpID := fakeLooper.LoopArgsForCall(0)

				Expect(actualID).To(Equal(providedCallID))
				Expect(actualScope).To(Equal(providedScope))
				Expect(actualSCG).To(Equal(providedSCG))
				Expect(actualOpHandle).To(Equal(providedOpHandle))
				Expect(actualParentID).To(Equal(providedParentID))
				Expect(actualRootOpID).To(Equal(providedRootOpID))
			})
		})
		It("should call callInterpreter.Interpret w/ expected args", func() {
			/* arrange */
			providedCallID := "dummyCallID"
			providedArgs := map[string]*model.Value{}
			providedSCG := &model.SCG{}
			providedOpHandle := new(data.FakeHandle)
			providedParentIDValue := "providedParentID"
			providedParentID := &providedParentIDValue
			providedRootOpID := "dummyRootOpID"

			fakeCallInterpreter := new(call.FakeInterpreter)
			fakeCallInterpreter.InterpretReturns(
				&model.DCG{},
				nil,
			)

			fakePubSub := new(pubsub.Fake)
			// ensure eventChan closed so call exits
			fakePubSub.SubscribeReturns(closedEventChan, nil)

			objectUnderTest := _caller{
				callInterpreter: fakeCallInterpreter,
				callStore:       new(fakeCallStore),
				containerCaller: new(fakeContainerCaller),
				pubSub:          fakePubSub,
			}

			/* act */
			objectUnderTest.Call(
				context.Background(),
				providedCallID,
				providedArgs,
				providedSCG,
				providedOpHandle,
				providedParentID,
				providedRootOpID,
			)

			/* assert */
			actualScope,
				actualSCG,
				actualID,
				actualOpHandle,
				actualParentID,
				actualRootOpID := fakeCallInterpreter.InterpretArgsForCall(0)

			Expect(actualScope).To(Equal(providedArgs))
			Expect(actualSCG).To(Equal(providedSCG))
			Expect(actualID).To(Equal(providedCallID))
			Expect(actualOpHandle).To(Equal(providedOpHandle))
			Expect(actualParentID).To(Equal(providedParentID))
			Expect(actualRootOpID).To(Equal(providedRootOpID))
		})
		Context("callInterpreter.Interpret result.If falsy", func() {
			It("should call pubSub.Publish w/ expected args", func() {
				/* arrange */
				providedCallID := "dummyCallID"
				providedRootOpID := "dummyRootOpID"

				fakeCallInterpreter := new(call.FakeInterpreter)
				falseBoolean := false
				fakeCallInterpreter.InterpretReturns(
					&model.DCG{
						If: &falseBoolean,
					},
					nil,
				)

				expectedEvent := model.Event{
					CallEnded: &model.CallEndedEvent{
						CallID:     providedCallID,
						RootCallID: providedRootOpID,
					},
				}

				fakePubSub := new(pubsub.Fake)
				// ensure eventChan closed so call exits
				fakePubSub.SubscribeReturns(closedEventChan, nil)

				objectUnderTest := _caller{
					callInterpreter: fakeCallInterpreter,
					callStore:       new(fakeCallStore),
					containerCaller: new(fakeContainerCaller),
					pubSub:          fakePubSub,
				}

				/* act */
				objectUnderTest.Call(
					context.Background(),
					providedCallID,
					map[string]*model.Value{},
					&model.SCG{},
					new(data.FakeHandle),
					nil,
					providedRootOpID,
				)

				/* assert */
				actualEvent := fakePubSub.PublishArgsForCall(0)

				// @TODO: implement/use VTime (similar to IOS & VFS) so we don't need custom assertions on temporal fields
				Expect(actualEvent.Timestamp).To(BeTemporally("~", time.Now().UTC(), 5*time.Second))
				// set temporal fields to expected vals since they're already asserted
				actualEvent.Timestamp = expectedEvent.Timestamp

				Expect(actualEvent).To(Equal(expectedEvent))
			})
		})

		Context("Container SCG", func() {
			It("should call containerCaller.Call w/ expected args", func() {
				/* arrange */
				fakeContainerCaller := new(fakeContainerCaller)

				providedArgs := map[string]*model.Value{}
				providedSCG := &model.SCG{
					Container: &model.SCGContainerCall{},
				}

				expectedDCG := &model.DCG{
					Container: &model.DCGContainerCall{},
				}
				fakeCallInterpreter := new(call.FakeInterpreter)
				fakeCallInterpreter.InterpretReturns(expectedDCG, nil)

				fakePubSub := new(pubsub.Fake)
				// ensure eventChan closed so call exits
				fakePubSub.SubscribeReturns(closedEventChan, nil)

				objectUnderTest := _caller{
					callInterpreter: fakeCallInterpreter,
					callStore:       new(fakeCallStore),
					containerCaller: fakeContainerCaller,
					pubSub:          fakePubSub,
				}

				/* act */
				objectUnderTest.Call(
					context.Background(),
					"dummyCallID",
					providedArgs,
					providedSCG,
					new(data.FakeHandle),
					nil,
					"dummyRootOpID",
				)

				/* assert */
				_,
					actualDCGContainerCall,
					actualArgs,
					actualSCG := fakeContainerCaller.CallArgsForCall(0)

				Expect(actualDCGContainerCall).To(Equal(expectedDCG.Container))
				Expect(actualArgs).To(Equal(providedArgs))
				Expect(actualSCG).To(Equal(providedSCG.Container))
			})
		})

		Context("Op SCG", func() {
			It("should call opCaller.Call w/ expected args", func() {
				/* arrange */
				fakeOpCaller := new(fakeOpCaller)

				expectedDCG := &model.DCG{
					Op: &model.DCGOpCall{},
				}
				fakeCallInterpreter := new(call.FakeInterpreter)
				fakeCallInterpreter.InterpretReturns(
					expectedDCG,
					nil,
				)

				providedCallID := "dummyCallID"
				providedArgs := map[string]*model.Value{}
				providedSCG := &model.SCG{
					Op: &model.SCGOpCall{
						Ref: "dummyOpRef",
					},
				}
				providedOpHandle := new(data.FakeHandle)
				providedParentID := "providedParentID"
				providedRootOpID := "dummyRootOpID"

				fakePubSub := new(pubsub.Fake)
				// ensure eventChan closed so call exits
				fakePubSub.SubscribeReturns(closedEventChan, nil)

				objectUnderTest := _caller{
					callInterpreter: fakeCallInterpreter,
					callStore:       new(fakeCallStore),
					opCaller:        fakeOpCaller,
					pubSub:          fakePubSub,
				}

				/* act */
				objectUnderTest.Call(
					context.Background(),
					providedCallID,
					providedArgs,
					providedSCG,
					providedOpHandle,
					&providedParentID,
					providedRootOpID,
				)

				/* assert */
				_,
					actualDCGOpCall,
					actualArgs,
					actualParentID,
					actualSCG := fakeOpCaller.CallArgsForCall(0)

				Expect(actualDCGOpCall).To(Equal(expectedDCG.Op))
				Expect(actualArgs).To(Equal(providedArgs))
				Expect(actualParentID).To(Equal(actualParentID))
				Expect(actualSCG).To(Equal(providedSCG.Op))
			})
		})

		Context("Parallel SCG", func() {
			It("should call parallelCaller.Call w/ expected args", func() {
				/* arrange */
				fakeParallelCaller := new(fakeParallelCaller)

				providedCallID := "dummyCallID"
				providedArgs := map[string]*model.Value{}
				providedSCG := &model.SCG{
					Parallel: []*model.SCG{
						{Container: &model.SCGContainerCall{}},
					},
				}
				providedOpHandle := new(data.FakeHandle)
				providedRootOpID := "dummyRootOpID"

				fakeCallInterpreter := new(call.FakeInterpreter)
				fakeCallInterpreter.InterpretReturns(
					&model.DCG{},
					nil,
				)

				fakePubSub := new(pubsub.Fake)
				// ensure eventChan closed so call exits
				fakePubSub.SubscribeReturns(closedEventChan, nil)

				objectUnderTest := _caller{
					callInterpreter: fakeCallInterpreter,
					callStore:       new(fakeCallStore),
					parallelCaller:  fakeParallelCaller,
					pubSub:          fakePubSub,
				}

				/* act */
				objectUnderTest.Call(
					context.Background(),
					providedCallID,
					providedArgs,
					providedSCG,
					providedOpHandle,
					nil,
					providedRootOpID,
				)

				/* assert */
				_,
					actualCallID,
					actualArgs,
					actualRootOpID,
					actualOpHandle,
					actualSCG := fakeParallelCaller.CallArgsForCall(0)

				Expect(actualCallID).To(Equal(providedCallID))
				Expect(actualArgs).To(Equal(providedArgs))
				Expect(actualRootOpID).To(Equal(providedRootOpID))
				Expect(actualOpHandle).To(Equal(providedOpHandle))
				Expect(actualSCG).To(Equal(providedSCG.Parallel))
			})
		})

		Context("Serial SCG", func() {

			It("should call serialCaller.Call w/ expected args", func() {
				/* arrange */
				fakeSerialCaller := new(fakeSerialCaller)

				providedCallID := "dummyCallID"
				providedArgs := map[string]*model.Value{}
				providedSCG := &model.SCG{
					Serial: []*model.SCG{
						{Container: &model.SCGContainerCall{}},
					},
				}
				providedOpHandle := new(data.FakeHandle)
				providedRootOpID := "dummyRootOpID"

				fakeCallInterpreter := new(call.FakeInterpreter)
				fakeCallInterpreter.InterpretReturns(
					&model.DCG{},
					nil,
				)

				fakePubSub := new(pubsub.Fake)
				// ensure eventChan closed so call exits
				fakePubSub.SubscribeReturns(closedEventChan, nil)

				objectUnderTest := _caller{
					callInterpreter: fakeCallInterpreter,
					callStore:       new(fakeCallStore),
					containerCaller: new(fakeContainerCaller),
					pubSub:          fakePubSub,
					serialCaller:    fakeSerialCaller,
				}

				/* act */
				objectUnderTest.Call(
					context.Background(),
					providedCallID,
					providedArgs,
					providedSCG,
					providedOpHandle,
					nil,
					providedRootOpID,
				)

				/* assert */
				_,
					actualCallID,
					actualArgs,
					actualRootOpID,
					actualOpHandle,
					actualSCG := fakeSerialCaller.CallArgsForCall(0)

				Expect(actualCallID).To(Equal(providedCallID))
				Expect(actualArgs).To(Equal(providedArgs))
				Expect(actualRootOpID).To(Equal(providedRootOpID))
				Expect(actualOpHandle).To(Equal(providedOpHandle))
				Expect(actualSCG).To(Equal(providedSCG.Serial))
			})
		})

		Context("No SCG", func() {
			It("should error", func() {
				/* arrange */
				providedCallID := "dummyCallID"
				providedArgs := map[string]*model.Value{}
				providedSCG := &model.SCG{}
				providedOpHandle := new(data.FakeHandle)
				providedRootOpID := "dummyRootOpID"
				expectedError := fmt.Errorf("Invalid call graph %+v\n", providedSCG)

				fakeCallInterpreter := new(call.FakeInterpreter)
				fakeCallInterpreter.InterpretReturns(
					&model.DCG{},
					nil,
				)

				fakePubSub := new(pubsub.Fake)
				// ensure eventChan closed so call exits
				fakePubSub.SubscribeReturns(closedEventChan, nil)

				objectUnderTest := _caller{
					callInterpreter: fakeCallInterpreter,
					callStore:       new(fakeCallStore),
					pubSub:          fakePubSub,
				}

				/* act */
				objectUnderTest.Call(
					context.Background(),
					providedCallID,
					providedArgs,
					providedSCG,
					providedOpHandle,
					nil,
					providedRootOpID,
				)

				/* assert */
				actualEvent := fakePubSub.PublishArgsForCall(0)

				Expect(actualEvent.CallEnded.Error.Message).To(Equal(expectedError.Error()))
			})
		})
	})
})
