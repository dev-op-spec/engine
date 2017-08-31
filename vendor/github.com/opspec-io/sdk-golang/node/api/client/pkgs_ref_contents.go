package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/node/api"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (c client) ListPkgContents(
	ctx context.Context,
	req model.ListPkgContentsReq,
) (
	[]*model.PkgContent,
	error,
) {

	// build url
	path := strings.Replace(api.URLPkgs_Ref_Contents, "{ref}", url.PathEscape(req.PkgRef), 1)

	httpReq, err := http.NewRequest(
		"GET",
		c.baseUrl.String()+path,
		nil,
	)
	if nil != err {
		return nil, err
	}

	httpReq = httpReq.WithContext(ctx)

	httpResp, err := c.httpClient.Do(httpReq)
	if nil != err {
		return nil, err
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode >= 400 {
		switch httpResp.StatusCode {
		case http.StatusUnauthorized:
			return nil, model.ErrPkgPullAuthentication{}
		case http.StatusForbidden:
			return nil, model.ErrPkgPullAuthorization{}
		case http.StatusNotFound:
			return nil, model.ErrPkgNotFound{}
		default:
			body, err := ioutil.ReadAll(httpResp.Body)
			if nil != err {
				return nil, fmt.Errorf(
					"Error encountered parsing response w/ status code '%v'; error was %v",
					httpResp.StatusCode,
					err.Error(),
				)
			}
			return nil, errors.New(string(body))
		}
	}

	var contentList []*model.PkgContent
	return contentList, json.NewDecoder(httpResp.Body).Decode(&contentList)

}
