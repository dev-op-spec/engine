package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/opctl/util/containerprovider"
	"github.com/opspec-io/opctl/util/pubsub"
	"github.com/opspec-io/opctl/util/uniquestring"
	"github.com/opspec-io/sdk-golang/pkg/model"
	"time"
)

var _ = Context("core", func() {
	Context("StartOp", func() {
		It("should call opCaller.Call w/ expected args", func() {
			/* arrange */
			providedReq := model.StartOpReq{
				Args: map[string]*model.Data{
					"dummyArg1Name": {String: "dummyArg1Value"},
					"dummyArg2Name": {Dir: "dummyArg2Value"},
					"dummyArg3Name": {Dir: "dummyArg3Value"},
					"dummyArg4Name": {Dir: "dummyArg4Value"},
				},
				PkgRef: "dummyPkgRef",
			}

			expectedOpId := "dummyOpId"

			fakeOpCaller := new(fakeOpCaller)

			fakeUniqueStringFactory := new(uniquestring.Fake)
			fakeUniqueStringFactory.ConstructReturns(expectedOpId)

			objectUnderTest := _core{
				containerProvider:   new(containerprovider.Fake),
				pubSub:              new(pubsub.Fake),
				opCaller:            fakeOpCaller,
				dcgNodeRepo:         new(fakeDcgNodeRepo),
				uniqueStringFactory: fakeUniqueStringFactory,
			}

			/* act */
			objectUnderTest.StartOp(providedReq)

			/* assert */
			// Call happens in go routine; wait 500ms to allow it to occur
			time.Sleep(time.Millisecond * 500)
			actualInputs,
				_,
				actualOpId,
				actualPkgRef,
				actualRootOpId := fakeOpCaller.CallArgsForCall(0)

			actualInputMap := map[string]*model.Data{}
			for actualInput := range actualInputs {
				actualInputMap[actualInput.Name] = actualInput.Value
			}

			Expect(actualInputMap).To(Equal(providedReq.Args))
			Expect(actualOpId).To(Equal(expectedOpId))
			Expect(actualPkgRef).To(Equal(providedReq.PkgRef))
			Expect(actualRootOpId).To(Equal(actualOpId))
		})
	})
})
