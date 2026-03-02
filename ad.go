package tiktokads

import (
	"errors"
	"time"
)

type createAdRequest struct {
	AccountId string `json:"advertiser_id"`
	AdGroupId string `json:"adgroup_id"`
	Creatives []*Ad  `json:"creatives"`
}

type createAdResponse struct {
	AdIds []string `json:"ad_ids"`
}

type Ad struct {
	Id                      string            `json:"ad_id,omitempty"`
	Name                    string            `json:"ad_name,omitempty"`
	IdentityType            string            `json:"identity_type,omitempty"`
	IdentityId              string            `json:"identity_id,omitempty"`
	CatalogId               string            `json:"catalog_id,omitempty"`
	ProductSetId            string            `json:"product_set_id,omitempty"`
	ProductSpecificType     string            `json:"product_specific_type,omitempty"`
	AdFormat                string            `json:"ad_format,omitempty"`
	AdText                  string            `json:"ad_text,omitempty"`
	DynamicFormat           AdDynamicFormat   `json:"dynamic_format,omitempty"`
	VerticalVideoStrategy   string            `json:"vertical_video_strategy,omitempty"`
	DynamicDestination      string            `json:"dynamic_destination"`
	CTA                     AdCta             `json:"call_to_action,omitempty"`
	AigcDisclosureType      string            `json:"aigc_disclosure_type,omitempty"`
	ShoppingAdsFallbackType string            `json:"shopping_ads_fallback_type,omitempty"`
	OperationStatus         AdOperationStatus `json:"operation_status,omitempty"`
	SecondaryStatus         string            `json:"secondary_status,omitempty"`
}

func (a *Ad) SetEnabled(enabled bool) {
	if enabled {
		a.OperationStatus = AdOperationStatus_ENABLE
	} else {
		a.OperationStatus = AdOperationStatus_DISABLE
	}
}

func (a *Ad) SetAdText(text string) {
	a.AdText = text
}

// SetProductSetId Set product set for ad. Set empty id to target all products
func (a *Ad) SetProductSetId(productSetId string) {
	if len(productSetId) == 0 {
		a.ProductSetId = ""
		a.ProductSpecificType = "ALL"
	} else {
		a.ProductSetId = productSetId
		a.ProductSpecificType = "PRODUCT_SET"
	}
}

func GetAds(accountId, adGroupId string) ([]*Ad, error) {
	req := newGetRequest(
		urlAdGet,
		withAccountId(accountId),
		withAdGroupIds([]string{adGroupId}),
	)

	return fetchAllPages[Ad](req, 100)
}

func UpdateAd(accountId, adGroupId string, ad *Ad) (*Ad, error) {
	adId := ad.Id
	if len(ad.Id) == 0 {
		req := newPostRequest(urlAdCreate, &createAdRequest{
			AccountId: accountId,
			AdGroupId: adGroupId,
			Creatives: []*Ad{ad},
		})

		created, err := fetch[createAdResponse](req)
		if nil != err {
			return nil, err
		}
		adId = created.AdIds[0]

		time.Sleep(4 * time.Second)
	} else if e := updateAdStatus(accountId, ad.Id, ad.OperationStatus); e != nil {
		return nil, e
	}

	ads, err := GetAds(accountId, adGroupId)
	if nil != err {
		return nil, err
	}

	for i, _ := range ads {
		if ads[i].Id == adId {
			return ads[i], nil
		}
	}

	return nil, errors.New("ad not found")
}

func DeleteAd(accountId, adId string) error {
	return updateAdStatus(accountId, adId, AdOperationStatus_DELETE)
}

func updateAdStatus(accountId, adId string, status AdOperationStatus) error {
	req := newPostRequest(
		urlAdStatusUpdate,
		&map[string]any{
			"advertiser_id":    accountId,
			"ad_ids":           []string{adId},
			"operation_status": status,
		},
	)

	_, err := fetch[emptyResult](req)

	return err

}
