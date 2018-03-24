package contents

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/golang-interfaces/encoding-ijson"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/data"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/node/core"
)

var _ = Context("Handler", func() {
	Context("Handle", func() {
		It("should call handle.ListDescendants", func() {
			/* arrange */
			fakeDataHandle := new(data.FakeHandle)
			// error to trigger immediate return
			fakeDataHandle.ListDescendantsReturns(nil, errors.New("dummyError"))

			fakeCore := new(core.Fake)
			fakeCore.ResolveDataReturns(fakeDataHandle, nil)

			objectUnderTest := _handler{
				core: fakeCore,
			}
			providedHTTPResp := httptest.NewRecorder()

			providedHTTPReq, err := http.NewRequest(
				http.MethodGet,
				"",
				nil,
			)
			if nil != err {
				panic(err.Error())
			}

			/* act */
			objectUnderTest.Handle(fakeDataHandle, providedHTTPResp, providedHTTPReq)

			/* assert */
			Expect(fakeDataHandle.ListDescendantsCallCount()).To(Equal(1))
		})
		Context("handle.ListDescendants errs", func() {
			It("should return expected result", func() {
				/* arrange */
				expectedBody := "dummyErrorMsg"

				fakeDataHandle := new(data.FakeHandle)
				// error to trigger immediate return
				fakeDataHandle.ListDescendantsReturns(nil, errors.New(expectedBody))

				fakeCore := new(core.Fake)
				fakeCore.ResolveDataReturns(fakeDataHandle, nil)

				objectUnderTest := _handler{
					core: fakeCore,
				}
				providedHTTPResp := httptest.NewRecorder()

				providedHTTPReq, err := http.NewRequest(
					http.MethodGet,
					"",
					nil,
				)
				if nil != err {
					panic(err.Error())
				}

				/* act */
				objectUnderTest.Handle(fakeDataHandle, providedHTTPResp, providedHTTPReq)

				/* assert */
				Expect(providedHTTPResp.Code).To(Equal(http.StatusInternalServerError))
				Expect(providedHTTPResp.HeaderMap.Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
				actualBody := strings.TrimSpace(providedHTTPResp.Body.String())
				Expect(actualBody).To(Equal(expectedBody))
			})
		})
		Context("handle.ListDescendants doesn't err", func() {
			Context("encoder.Encode errs", func() {
				It("should return expected result", func() {
					/* arrange */
					expectedBody := "dummyErrorMsg"

					fakeCore := new(core.Fake)

					fakeJSON := new(ijson.Fake)
					fakeJSON.NewEncoderReturns(json.NewEncoder(errWriter{Msg: expectedBody}))

					objectUnderTest := _handler{
						core: fakeCore,
						json: fakeJSON,
					}

					providedHTTPResp := httptest.NewRecorder()

					providedHTTPReq, err := http.NewRequest(
						http.MethodGet,
						"",
						nil,
					)
					if nil != err {
						panic(err.Error())
					}

					/* act */
					objectUnderTest.Handle(new(data.FakeHandle), providedHTTPResp, providedHTTPReq)

					/* assert */
					Expect(providedHTTPResp.Code).To(Equal(http.StatusInternalServerError))
					Expect(providedHTTPResp.HeaderMap.Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
					actualBody := strings.TrimSpace(providedHTTPResp.Body.String())
					Expect(actualBody).To(Equal(expectedBody))
				})
			})
			Context("encoder.Encode doesn't err", func() {
				It("should return expected result", func() {
					/* arrange */

					fakeCore := new(core.Fake)

					fakeHandle := new(data.FakeHandle)
					contentsList := []*model.DataNode{
						{Path: "dummyPath"},
					}
					fakeHandle.ListDescendantsReturns(contentsList, nil)

					expectedBodyBytes, err := json.Marshal(contentsList)
					if nil != err {
						panic(err)
					}

					fakeJSON := new(ijson.Fake)
					fakeJSON.NewEncoderStub = func(w io.Writer) *json.Encoder {
						return json.NewEncoder(w)
					}

					objectUnderTest := _handler{
						core: fakeCore,
						json: fakeJSON,
					}

					providedHTTPResp := httptest.NewRecorder()

					providedHTTPReq, err := http.NewRequest(
						http.MethodGet,
						"",
						nil,
					)
					if nil != err {
						panic(err.Error())
					}

					/* act */
					objectUnderTest.Handle(fakeHandle, providedHTTPResp, providedHTTPReq)

					/* assert */
					Expect(providedHTTPResp.Code).To(Equal(http.StatusOK))
					Expect(providedHTTPResp.HeaderMap.Get("Content-Type")).To(Equal("application/json; charset=UTF-8"))
					actualBody := strings.TrimSpace(providedHTTPResp.Body.String())
					Expect(actualBody).To(Equal(string(expectedBodyBytes)))
				})
			})
		})
	})
})

// errWriter always errs
type errWriter struct {
	Msg string
}

func (this errWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New(this.Msg)
}
