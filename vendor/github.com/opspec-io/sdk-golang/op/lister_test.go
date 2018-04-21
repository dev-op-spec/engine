package op

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/data"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/op/dotyml"
)

var _ = Context("Lister", func() {
	Context("NewLister", func() {
		It("should not return nil", func() {
			/* arrange/act/assert */
			Expect(NewLister()).Should(Not(BeNil()))
		})
	})
	Context("List", func() {
		It("should call dirHandle.ListDescendants w/ expected args", func() {
			/* arrange */
			providedCtx := context.Background()

			providedDirHandle := new(data.FakeHandle)
			// err to trigger immediate return
			providedDirHandle.ListDescendantsReturns(nil, errors.New("dummyError"))

			objectUnderTest := _lister{}

			/* act */
			objectUnderTest.List(
				providedCtx,
				providedDirHandle,
			)

			/* assert */
			actualCtx := providedDirHandle.ListDescendantsArgsForCall(0)

			Expect(actualCtx).To(Equal(providedCtx))
		})
		Context("dirHandle.ListDescendants errs", func() {
			It("should return expected result", func() {
				/* arrange */
				providedDirHandle := new(data.FakeHandle)
				listDescendantsErr := errors.New("listDescendantsErr")
				providedDirHandle.ListDescendantsReturns(nil, listDescendantsErr)

				objectUnderTest := _lister{}

				/* act */
				_, actualErr := objectUnderTest.List(
					context.Background(),
					providedDirHandle,
				)

				/* assert */
				Expect(actualErr).To(Equal(listDescendantsErr))
			})
		})
		Context("dirHandle.ListDescendants doesn't err", func() {
			Context("dirHandle.ListDescendants returns items", func() {
				Context("item.Path ends w/ op.yml", func() {
					It("should call dirHandle.GetContent w/ expected args", func() {
						/* arrange */
						providedCtx := context.Background()

						providedDirHandle := new(data.FakeHandle)
						item := model.DirEntry{
							Path: dotyml.FileName,
						}
						providedDirHandle.ListDescendantsReturns([]*model.DirEntry{&item}, nil)

						// err to trigger immediate return
						providedDirHandle.GetContentReturns(nil, errors.New("dummyError"))

						objectUnderTest := _lister{}

						/* act */
						objectUnderTest.List(
							providedCtx,
							providedDirHandle,
						)

						/* assert */
						actualCtx,
							actualPath := providedDirHandle.GetContentArgsForCall(0)

						Expect(actualCtx).To(Equal(providedCtx))
						Expect(actualPath).To(Equal(item.Path))

					})
					Context("dirHandle.GetContent errs", func() {
						It("should return expected result", func() {
							/* arrange */
							providedDirHandle := new(data.FakeHandle)
							providedDirHandle.ListDescendantsReturns([]*model.DirEntry{{}}, nil)

							getContentErr := errors.New("getContentErr")
							providedDirHandle.GetContentReturns(nil, getContentErr)

							objectUnderTest := _lister{}

							/* act */
							actualOpYmls, actualErr := objectUnderTest.List(
								context.Background(),
								providedDirHandle,
							)

							/* assert */
							Expect(actualOpYmls).To(BeEmpty())
							Expect(actualErr).To(BeNil())
						})
					})
				})
			})
		})
	})
})
