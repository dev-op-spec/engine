package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/opctl/util/containerprovider"
	"github.com/opspec-io/opctl/util/pubsub"
	"github.com/opspec-io/sdk-golang/pkg/model"
	"github.com/pkg/errors"
	"time"
)

var _ = Context("containerCaller", func() {
	Context("newContainerCaller", func() {
		It("should return containerCaller", func() {
			/* arrange/act/assert */
			Expect(newContainerCaller(
				new(containerprovider.Fake),
				new(pubsub.Fake),
				new(fakeDcgNodeRepo),
			)).Should(Not(BeNil()))
		})
	})
	Context("Call", func() {
		It("should call dcgNodeRepo.Add w/ expected args", func() {
			/* arrange */
			providedInputs := make(chan *variable)
			// inputs chan must be closed for method under test to return
			close(providedInputs)

			providedOutputs := make(chan *variable, 150)
			providedContainerId := "dummyContainerId"
			providedScgContainerCall := &model.ScgContainerCall{}
			providedPkgRef := "dummyPkgRef"
			providedRootOpId := "dummyRootOpId"

			fakePubSub := new(pubsub.Fake)
			fakePubSub.SubscribeStub = func(filter *model.EventFilter, eventChannel chan *model.Event) {
				close(eventChannel)
			}

			expectedDcgNodeDescriptor := &dcgNodeDescriptor{
				Id:        providedContainerId,
				PkgRef:    providedPkgRef,
				RootOpId:  providedRootOpId,
				Container: &dcgContainerDescriptor{},
			}

			fakeDcgNodeRepo := new(fakeDcgNodeRepo)

			objectUnderTest := newContainerCaller(
				new(containerprovider.Fake),
				fakePubSub,
				fakeDcgNodeRepo,
			)

			/* act */
			objectUnderTest.Call(
				providedInputs,
				providedOutputs,
				providedContainerId,
				providedScgContainerCall,
				providedPkgRef,
				providedRootOpId,
			)

			/* assert */
			Expect(fakeDcgNodeRepo.AddArgsForCall(0)).To(Equal(expectedDcgNodeDescriptor))
		})
		It("should call pubSub.Publish w/ expected args", func() {
			/* arrange */
			providedInputs := make(chan *variable)
			// inputs chan must be closed for method under test to return
			close(providedInputs)

			providedOutputs := make(chan *variable, 150)
			providedContainerId := "dummyContainerId"
			providedScgContainerCall := &model.ScgContainerCall{}
			providedPkgRef := "dummyPkgRef"
			providedRootOpId := "dummyRootOpId"

			expectedEvent := &model.Event{
				Timestamp: time.Now().UTC(),
				ContainerStarted: &model.ContainerStartedEvent{
					ContainerId: providedContainerId,
					PkgRef:      providedPkgRef,
					RootOpId:    providedRootOpId,
				},
			}

			fakePubSub := new(pubsub.Fake)
			fakePubSub.SubscribeStub = func(filter *model.EventFilter, eventChannel chan *model.Event) {
				close(eventChannel)
			}

			objectUnderTest := newContainerCaller(
				new(containerprovider.Fake),
				fakePubSub,
				new(fakeDcgNodeRepo),
			)

			/* act */
			objectUnderTest.Call(
				providedInputs,
				providedOutputs,
				providedContainerId,
				providedScgContainerCall,
				providedPkgRef,
				providedRootOpId,
			)

			/* assert */
			actualEvent := fakePubSub.PublishArgsForCall(0)

			// @TODO: implement/use VTime (similar to VOS & VFS) so we don't need custom assertions on temporal fields
			Expect(actualEvent.Timestamp).To(BeTemporally("~", time.Now().UTC(), 5*time.Second))
			// set temporal fields to expected vals since they're already asserted
			actualEvent.Timestamp = expectedEvent.Timestamp

			Expect(actualEvent).To(Equal(expectedEvent))
		})
		It("should call containerProvider.RunContainer w/ expected args", func() {
			/* arrange */
			providedInputs := make(chan *variable)
			// inputs chan must be closed for method under test to return
			close(providedInputs)

			providedOutputs := make(chan *variable, 150)
			providedContainerId := "dummyContainerId"
			providedScgContainerCall := &model.ScgContainerCall{}
			providedPkgRef := "dummyPkgRef"
			providedRootOpId := "dummyRootOpId"

			expectedReq, _ := constructDcgContainerCall(
				map[string]*model.Data{},
				providedScgContainerCall,
				providedContainerId,
				providedRootOpId,
				providedPkgRef,
			)

			fakeContainerProvider := new(containerprovider.Fake)

			fakePubSub := new(pubsub.Fake)
			fakePubSub.SubscribeStub = func(filter *model.EventFilter, eventChannel chan *model.Event) {
				close(eventChannel)
			}

			objectUnderTest := newContainerCaller(
				fakeContainerProvider,
				fakePubSub,
				new(fakeDcgNodeRepo),
			)

			/* act */
			objectUnderTest.Call(
				providedInputs,
				providedOutputs,
				providedContainerId,
				providedScgContainerCall,
				providedPkgRef,
				providedRootOpId,
			)

			/* assert */
			actualReq, actualEventPublisher := fakeContainerProvider.RunContainerArgsForCall(0)
			Expect(actualReq).To(Equal(expectedReq))
			Expect(actualEventPublisher).To(Equal(fakePubSub))
		})
		Context("containerProvider.RunContainer errors", func() {
			It("should return expected error", func() {
				/* arrange */
				providedInputs := make(chan *variable)
				// inputs chan must be closed for method under test to return
				close(providedInputs)

				providedOutputs := make(chan *variable, 150)
				providedContainerId := "dummyContainerId"
				providedScgContainerCall := &model.ScgContainerCall{}
				providedPkgRef := "dummyPkgRef"
				providedRootOpId := "dummyRootOpId"

				expectedError := errors.New("dummyError")

				fakeContainerProvider := new(containerprovider.Fake)
				fakeContainerProvider.RunContainerReturns(expectedError)

				objectUnderTest := newContainerCaller(
					fakeContainerProvider,
					new(pubsub.Fake),
					new(fakeDcgNodeRepo),
				)

				/* act */
				actualError := objectUnderTest.Call(
					providedInputs,
					providedOutputs,
					providedContainerId,
					providedScgContainerCall,
					providedPkgRef,
					providedRootOpId,
				)

				/* assert */
				Expect(actualError).To(Equal(expectedError))
			})
		})
	})
	It("should call dcgNodeRepo.DeleteIfExists w/ expected args", func() {
		/* arrange */
		providedInputs := make(chan *variable)
		// inputs chan must be closed for method under test to return
		close(providedInputs)

		providedOutputs := make(chan *variable, 150)
		providedContainerId := "dummyContainerId"
		providedScgContainerCall := &model.ScgContainerCall{}
		providedPkgRef := "dummyPkgRef"
		providedRootOpId := "dummyRootOpId"

		fakePubSub := new(pubsub.Fake)
		fakePubSub.SubscribeStub = func(filter *model.EventFilter, eventChannel chan *model.Event) {
			close(eventChannel)
		}

		fakeDcgNodeRepo := new(fakeDcgNodeRepo)

		objectUnderTest := newContainerCaller(
			new(containerprovider.Fake),
			fakePubSub,
			fakeDcgNodeRepo,
		)

		/* act */
		objectUnderTest.Call(
			providedInputs,
			providedOutputs,
			providedContainerId,
			providedScgContainerCall,
			providedPkgRef,
			providedRootOpId,
		)

		/* assert */
		Expect(fakeDcgNodeRepo.DeleteIfExistsArgsForCall(0)).To(Equal(providedContainerId))
	})
	It("should call pubSub.Publish w/ expected args", func() {
		/* arrange */
		providedInputs := make(chan *variable)
		// inputs chan must be closed for method under test to return
		close(providedInputs)

		providedOutputs := make(chan *variable, 150)
		providedContainerId := "dummyContainerId"
		providedScgContainerCall := &model.ScgContainerCall{}
		providedPkgRef := "dummyPkgRef"
		providedRootOpId := "dummyRootOpId"

		expectedEvent := &model.Event{
			Timestamp: time.Now().UTC(),
			ContainerExited: &model.ContainerExitedEvent{
				ContainerId: providedContainerId,
				PkgRef:      providedPkgRef,
				RootOpId:    providedRootOpId,
			},
		}

		fakePubSub := new(pubsub.Fake)
		fakePubSub.SubscribeStub = func(filter *model.EventFilter, eventChannel chan *model.Event) {
			close(eventChannel)
		}

		objectUnderTest := newContainerCaller(
			new(containerprovider.Fake),
			fakePubSub,
			new(fakeDcgNodeRepo),
		)

		/* act */
		objectUnderTest.Call(
			providedInputs,
			providedOutputs,
			providedContainerId,
			providedScgContainerCall,
			providedPkgRef,
			providedRootOpId,
		)

		/* assert */
		actualEvent := fakePubSub.PublishArgsForCall(1)

		// @TODO: implement/use VTime (similar to VOS & VFS) so we don't need custom assertions on temporal fields
		Expect(actualEvent.Timestamp).To(BeTemporally("~", time.Now().UTC(), 5*time.Second))
		// set temporal fields to expected vals since they're already asserted
		actualEvent.Timestamp = expectedEvent.Timestamp

		Expect(actualEvent).To(Equal(expectedEvent))
	})
})
