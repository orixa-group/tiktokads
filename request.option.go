package tiktokads

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type requestOption func(req *http.Request)

// withRequestToken Add access_token header for business api
func withRequestToken(req *http.Request) {
	req.Header.Set("Access-Token", accessToken)
}

func withJsonContentHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}

func withQueryStringArrayParam(name string, values []string) requestOption {
	buf, _ := json.Marshal(values)
	return withQueryString(map[string]string{
		name: string(buf),
	})
}

func withQueryString(params map[string]string) requestOption {
	return func(req *http.Request) {
		q := req.URL.Query()

		for k, v := range params {
			q.Set(k, v)
		}

		req.URL.RawQuery = q.Encode()
	}
}

// withPageNumber Set page number
func withPageNumber(pageNumber int) requestOption {
	return withQueryString(map[string]string{
		"page": strconv.Itoa(pageNumber),
	})
}

// withPerPage Set result per page
func withPerPage(perPage int) requestOption {
	return withQueryString(map[string]string{
		"page_size": strconv.Itoa(perPage),
	})
}

// withAppIdAndSecret Add app_id & secret to query string. Required for some apis.
func withAppIdAndSecret(req *http.Request) {
	withQueryString(map[string]string{
		"app_id": appId,
		"secret": appSecret,
	})(req)
}

// withAccountId Add advertiser_id to query string. Required for campaign apis
func withAccountId(id string) requestOption {
	return withQueryString(map[string]string{
		"advertiser_id": id,
	})
}

// withBusinessId Add bc_id to query string. Required for catalog creation api
func withBusinessId(id string) requestOption {
	return withQueryString(map[string]string{
		"bc_id": id,
	})
}

// withCatalogId Add catalog_id to query string. Required for catalog creation api
func withCatalogId(id string) requestOption {
	return withQueryString(map[string]string{
		"catalog_id": id,
	})
}

// withFiltering Add filters as json encoded object to filtering params.
func withFiltering(filters map[string]any) requestOption {
	buf, _ := json.Marshal(&filters)
	return withQueryString(map[string]string{
		"filtering": string(buf),
	})
}
