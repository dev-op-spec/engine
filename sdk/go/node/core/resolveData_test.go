package core

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/sdk-golang/data"
	"github.com/opctl/sdk-golang/model"
)

var _ = Context("core", func() {
	Context("ResolveData", func() {
		It("should call data.Resolve w/ expected args", func() {
			/* arrange */
			providedCtx := context.Background()
			providedOpRef := "dummyOpRef"
			providedPullCreds := &model.PullCreds{
				Username: "dummyUsername",
				Password: "dummyPassword",
			}

			fakeData := new(data.Fake)

			expectedPkgProviders := []data.Provider{
				new(data.FakeProvider),
				new(data.FakeProvider),
			}
			fakeData.NewFSProviderReturns(expectedPkgProviders[0])
			fakeData.NewGitProviderReturns(expectedPkgProviders[1])

			objectUnderTest := _core{
				data: fakeData,
			}

			/* act */
			objectUnderTest.ResolveData(
				providedCtx,
				providedOpRef,
				providedPullCreds,
			)

			/* assert */
			actualCtx,
				actualOpRef,
				actualPkgProviders := fakeData.ResolveArgsForCall(0)

			Expect(actualCtx).To(Equal(providedCtx))
			Expect(actualOpRef).To(Equal(providedOpRef))
			Expect(actualPkgProviders).To(ConsistOf(expectedPkgProviders))
		})
	})
})
