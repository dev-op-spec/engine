package data

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/node/api/client"
)

var _ = Context("fsHandle", func() {

	Context("GetContent", func() {

		It("should call client.GetPkgContent w/ expected args", func() {
			/* arrange */
			providedCtx := context.TODO()
			providedContentPath := "dummyContentPath"

			dataRef := "dummyPkgRef"
			pullCreds := &model.PullCreds{Username: "dummyUsername", Password: "dummyPassword"}

			fakeClient := new(client.Fake)

			objectUnderTest := nodeHandle{
				client:    fakeClient,
				dataRef:   dataRef,
				pullCreds: pullCreds,
			}

			/* act */
			objectUnderTest.GetContent(providedCtx, providedContentPath)

			/* assert */
			actualCtx,
				actualReq := fakeClient.GetPkgContentArgsForCall(0)

			Expect(actualCtx).To(Equal(providedCtx))
			Expect(actualReq).To(Equal(model.GetPkgContentReq{
				ContentPath: providedContentPath,
				PkgRef:      dataRef,
				PullCreds:   pullCreds,
			}))
		})
	})

	Context("ListContents", func() {
		It("should call client.ListPkgContents w/ expected args", func() {
			/* arrange */
			providedCtx := context.TODO()

			dataRef := "dummyPkgRef"
			pullCreds := &model.PullCreds{Username: "dummyUsername", Password: "dummyPassword"}

			fakeClient := new(client.Fake)

			objectUnderTest := nodeHandle{
				client:    fakeClient,
				dataRef:   dataRef,
				pullCreds: pullCreds,
			}

			/* act */
			objectUnderTest.ListContents(providedCtx)

			/* assert */
			actualCtx,
				actualReq := fakeClient.ListPkgContentsArgsForCall(0)

			Expect(actualCtx).To(Equal(providedCtx))
			Expect(actualReq).To(Equal(model.ListPkgContentsReq{
				PkgRef:    dataRef,
				PullCreds: pullCreds,
			}))
		})
	})

	Context("Ref", func() {
		It("should return expected ref", func() {
			/* arrange */
			dataRef := "dummyPkgRef"

			fakeClient := new(client.Fake)

			objectUnderTest := nodeHandle{
				client:  fakeClient,
				dataRef: dataRef,
			}

			/* act */
			actualRef := objectUnderTest.Ref()

			/* assert */
			Expect(actualRef).To(Equal(dataRef))
		})
	})
})