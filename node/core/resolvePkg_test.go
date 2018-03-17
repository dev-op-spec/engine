package core

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/data"
	"github.com/opspec-io/sdk-golang/model"
)

var _ = Context("core", func() {
	Context("ResolvePkg", func() {
		It("should call data.Resolve w/ expected args", func() {
			/* arrange */
			providedCtx := context.Background()
			providedPkgRef := "dummyPkgRef"
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
			objectUnderTest.ResolvePkg(
				providedCtx,
				providedPkgRef,
				providedPullCreds,
			)

			/* assert */
			actualCtx,
				actualPkgRef,
				actualPkgProviders := fakeData.ResolveArgsForCall(0)

			Expect(actualCtx).To(Equal(providedCtx))
			Expect(actualPkgRef).To(Equal(providedPkgRef))
			Expect(actualPkgProviders).To(ConsistOf(expectedPkgProviders))
		})
	})
})
