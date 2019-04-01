package core

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/sdk-golang/model"
	"github.com/opctl/sdk-golang/util/pubsub"
)

var _ = Context("core", func() {
	Context("GetEventStream", func() {
		It("should call pubSub.Subscribe w/ expected args", func() {
			/* arrange */
			providedCtx := context.TODO()
			providedReq := &model.GetEventStreamReq{
				Filter: model.EventFilter{
					Roots: []string{
						"dummyRootOpID",
					},
				},
			}

			fakePubSub := new(pubsub.Fake)

			objectUnderTest := _core{
				pubSub: fakePubSub,
			}

			/* act */
			objectUnderTest.GetEventStream(
				providedCtx,
				providedReq,
			)

			/* assert */

			actualCtx,
				actualFilter := fakePubSub.SubscribeArgsForCall(0)

			Expect(actualCtx).To(Equal(providedCtx))
			Expect(actualFilter).To(Equal(providedReq.Filter))
		})
	})
})
