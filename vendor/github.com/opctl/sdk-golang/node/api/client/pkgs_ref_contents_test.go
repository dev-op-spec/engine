package client

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang-interfaces/ihttp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/sdk-golang/model"
	"github.com/opctl/sdk-golang/node/api"
)

var _ = Context("ListDescendants", func() {

	It("should call httpClient.Do() w/ expected args & return result", func() {

		/* arrange */
		providedCtx := context.TODO()
		providedReq := model.ListDescendantsReq{
			PkgRef: "dummyOpRef",
			PullCreds: &model.PullCreds{
				Username: "dummyUsername",
				Password: "dummyPassword",
			},
		}

		expectedReqURL := url.URL{}
		path := strings.Replace(api.URLPkgs_Ref_Contents, "{ref}", url.PathEscape(providedReq.PkgRef), 1)
		expectedReqURL.Path = path

		expectedHTTPReq, _ := http.NewRequest(
			"GET",
			expectedReqURL.String(),
			nil,
		)

		expectedHTTPReq.SetBasicAuth(
			providedReq.PullCreds.Username,
			providedReq.PullCreds.Password,
		)

		fakeHttpClient := new(ihttp.FakeClient)
		fakeHttpClient.DoReturns(
			&http.Response{
				Body:       ioutil.NopCloser(strings.NewReader("[]")),
				StatusCode: http.StatusOK,
			},
			nil,
		)

		objectUnderTest := client{
			httpClient: fakeHttpClient,
		}

		/* act */
		objectUnderTest.ListDescendants(providedCtx, providedReq)

		/* assert */
		actualHTTPReq := fakeHttpClient.DoArgsForCall(0)

		Expect(actualHTTPReq.URL).To(Equal(expectedHTTPReq.URL))
		Expect(actualHTTPReq.Body).To(BeNil())
		Expect(actualHTTPReq.Header).To(Equal(expectedHTTPReq.Header))
		Expect(actualHTTPReq.Context()).To(Equal(providedCtx))

	})
	Context("StatusCode < 400", func() {

		It("should return expected result", func() {

			/* arrange */
			httpResp := &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader("[]")),
				StatusCode: http.StatusOK,
			}

			fakeHttpClient := new(ihttp.FakeClient)
			fakeHttpClient.DoReturns(httpResp, nil)

			objectUnderTest := client{
				httpClient: fakeHttpClient,
			}

			/* act */
			actualContentsList, actualErr := objectUnderTest.ListDescendants(
				context.TODO(),
				model.ListDescendantsReq{},
			)

			/* assert */
			Expect(actualContentsList).To(Equal([]*model.DirEntry{}))
			Expect(actualErr).To(BeNil())

		})
	})
	Context("StatusCode >= 400", func() {
		Context("401", func() {
			It("should return expected result", func() {

				/* arrange */
				httpResp := &http.Response{
					Body:       ioutil.NopCloser(nil),
					StatusCode: http.StatusUnauthorized,
				}

				fakeHttpClient := new(ihttp.FakeClient)
				fakeHttpClient.DoReturns(httpResp, nil)

				objectUnderTest := client{
					httpClient: fakeHttpClient,
				}

				/* act */
				_, actualErr := objectUnderTest.ListDescendants(
					context.TODO(),
					model.ListDescendantsReq{},
				)

				/* assert */
				Expect(actualErr).To(Equal(model.ErrDataProviderAuthentication{}))

			})
		})
		Context("403", func() {
			It("should return expected result", func() {

				/* arrange */
				httpResp := &http.Response{
					Body:       ioutil.NopCloser(nil),
					StatusCode: http.StatusForbidden,
				}

				fakeHttpClient := new(ihttp.FakeClient)
				fakeHttpClient.DoReturns(httpResp, nil)

				objectUnderTest := client{
					httpClient: fakeHttpClient,
				}

				/* act */
				_, actualErr := objectUnderTest.ListDescendants(
					context.TODO(),
					model.ListDescendantsReq{},
				)

				/* assert */
				Expect(actualErr).To(Equal(model.ErrDataProviderAuthorization{}))

			})

		})
		Context("404", func() {
			It("should return expected result", func() {

				/* arrange */
				httpResp := &http.Response{
					Body:       ioutil.NopCloser(nil),
					StatusCode: http.StatusNotFound,
				}

				fakeHttpClient := new(ihttp.FakeClient)
				fakeHttpClient.DoReturns(httpResp, nil)

				objectUnderTest := client{
					httpClient: fakeHttpClient,
				}

				/* act */
				_, actualErr := objectUnderTest.ListDescendants(
					context.TODO(),
					model.ListDescendantsReq{},
				)

				/* assert */
				Expect(actualErr).To(Equal(model.ErrDataRefResolution{}))

			})

		})
		Context("500", func() {
			It("should return expected result", func() {

				/* arrange */
				expectedErr := errors.New("dummyMsg")
				httpResp := &http.Response{
					Body:       ioutil.NopCloser(strings.NewReader(expectedErr.Error())),
					StatusCode: http.StatusInternalServerError,
				}

				fakeHttpClient := new(ihttp.FakeClient)
				fakeHttpClient.DoReturns(httpResp, nil)

				objectUnderTest := client{
					httpClient: fakeHttpClient,
				}

				/* act */
				_, actualErr := objectUnderTest.ListDescendants(
					context.TODO(),
					model.ListDescendantsReq{},
				)

				/* assert */
				Expect(actualErr).To(Equal(expectedErr))

			})

		})
	})
})
