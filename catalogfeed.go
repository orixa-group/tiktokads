package tiktokads

import (
	"errors"
	"net/http"
	"time"
)

type catalogFeedParam struct {
	Uri           string `json:"uri"`
	UpdateMode    string `json:"update_mode"`
	Timezone      string `json:"timezone"`
	IntervalType  string `json:"interval_type"`
	IntervalCount int    `json:"interval_count"`
}

type CatalogFeed struct {
	FeedId          string            `json:"feed_id"`
	FeedName        string            `json:"feed_name"`
	LastUpdateParam *catalogFeedParam `json:"last_update_param"`
	// Read-only fields
	Status   string `json:"status"`
	Products int    `json:"number_of_products"`
}

func NewCatalogFeed(feedName, url string) *CatalogFeed {
	return &CatalogFeed{
		FeedName: feedName,
		LastUpdateParam: &catalogFeedParam{
			Uri:           url,
			Timezone:      "Europe/Paris",
			IntervalType:  "HOURLY",
			IntervalCount: 1,
		},
	}
}

type catalogFeedUpdate struct {
	// not set for creation
	Id         string         `json:"feed_id,omitempty"`
	BcId       string         `json:"bc_id,omitempty"`
	CatalogId  string         `json:"catalog_id,omitempty"`
	UpdateMode string         `json:"update_mode,omitempty"`
	FeedName   string         `json:"feed_name,omitempty"`
	Schedule   *scheduleParam `json:"schedule_param,omitempty"`
}

type scheduleParam struct {
	IntervalType  string               `json:"interval_type,omitempty"`
	IntervalCount int                  `json:"interval_count,omitempty"`
	Minute        *int                 `json:"minute,omitempty"`
	Timezone      string               `json:"timezone,omitempty"`
	Source        *scheduleParamSource `json:"source,omitempty"`
}

type scheduleParamSource struct {
	Uri string `json:"uri,omitempty"`
}

func GetCatalogFeeds(businessId, catalogId string) ([]*CatalogFeed, error) {
	req := newGetRequest(
		urlCatalogFeedGet,
		withBusinessId(businessId),
		withCatalogId(catalogId),
	)

	feeds, err := fetch[listFeedResult[CatalogFeed]](req)

	if nil != err {
		return nil, err
	}

	return feeds.FeedList, err
}

func newCatalogFeedCreate(catalog *CatalogFeed) *catalogFeedUpdate {
	// set schedule to next minute for launch feed download on creation
	nextMinute := time.Now().Minute() + 2

	return &catalogFeedUpdate{
		Id:         catalog.FeedId,
		UpdateMode: "OVERWRITE",
		FeedName:   catalog.FeedName,
		Schedule: &scheduleParam{
			IntervalCount: catalog.LastUpdateParam.IntervalCount,
			IntervalType:  catalog.LastUpdateParam.IntervalType,
			Timezone:      catalog.LastUpdateParam.Timezone,
			Minute:        &nextMinute,
			Source: &scheduleParamSource{
				Uri: catalog.LastUpdateParam.Uri,
			},
		},
	}
}

func UpdateCatalogFeed(businessId, catalogId string, feed *CatalogFeed) (*CatalogFeed, error) {
	body := newCatalogFeedCreate(feed)
	body.CatalogId = catalogId
	body.BcId = businessId
	var req *http.Request

	type response struct {
		FeedId string `json:"feed_id"`
	}

	if len(body.Id) > 0 {
		req = newPostRequest(urlCatalogFeedUpdate, body)
	} else {
		req = newPostRequest(urlCatalogFeedCreate, body)
	}

	if resp, err := fetch[response](req); nil != err {
		return nil, err
	} else {

		if feeds, err := GetCatalogFeeds(businessId, catalogId); nil != err {
			return nil, err
		} else {
			for _, f := range feeds {
				if f.FeedId == resp.FeedId {
					return f, nil
				}
			}
		}
	}

	return nil, errors.New("feed not found")
}
