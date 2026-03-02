package tiktokads

import (
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var LastResponse = ""
var LastRequestPayload = ""
var cli = &http.Client{
	Timeout: time.Second * 30,
}

// newGetRequest Builder for GET requests
func newGetRequest(apiUrl string, options ...requestOption) *http.Request {
	req, e := http.NewRequest(http.MethodGet, apiUrl, nil)
	if nil != e {
		panic(e)
	}

	for _, o := range options {
		o(req)
	}

	return req
}

// newPostRequest Builder for POST requests with json body.
func newPostRequest[T any](apiUrl string, body *T, options ...requestOption) *http.Request {
	var r io.Reader

	if body != nil {
		buf, _ := json.Marshal(body)
		LastRequestPayload = string(buf)
		r = bytes.NewReader(buf)
	}

	req, e := http.NewRequest(http.MethodPost, apiUrl, r)
	if nil != e {
		panic(e)
	}
	for _, o := range options {
		o(req)
	}
	withJsonContentHeader(req)

	return req
}

// fetch Fetch response from request. Request can either be GET or POST.
func fetch[T any](req *http.Request) (*T, error) {
	withRequestToken(req)

	resp, err := cli.Do(req)

	if err != nil {
		return nil, err
	} else if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("tiktok ads api error: %s", resp.Status)
	}

	var res apiResult[T]
	var buf []byte
	buf, err = io.ReadAll(resp.Body)

	if err = cmp.Or(
		json.Unmarshal(buf, &res),
		err,
	); nil != err {
		return nil, err
	} else if res.Code > 0 {
		// check body for error because api doesn't always return http 4xx/5xx error codes
		return nil, fmt.Errorf("%s (%d)", res.Message, res.Code)
	}

	return res.Data, nil
}

// fetchAllPages Fetch all response from request with paginated results.
// perPage limit param depends on api objects
func fetchAllPages[T any](req *http.Request, perPage int) ([]*T, error) {
	withPerPage(perPage)(req)

	var allResults []*T

	for i := 1; ; i++ {
		withPageNumber(i)(req)
		if res, err := fetch[listResult[T]](req); nil != err {
			return nil, err
		} else {
			allResults = append(allResults, res.GetResults()...)
			if len(res.List) < perPage {
				break
			}
		}
	}

	return allResults, nil
}
