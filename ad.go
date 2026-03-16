package tiktokads

import (
	"errors"
	"reflect"
	"time"
)

type createAdRequest struct {
	AccountId   string `json:"advertiser_id"`
	AdGroupId   string `json:"adgroup_id"`
	Creatives   []*Ad  `json:"creatives"`
	PatchUpdate *bool  `json:"patch_update,omitempty"` // for updates
}

type createAdResponse struct {
	AdIds []string `json:"ad_ids"`
}

type Ad struct {
	Id                      string            `json:"ad_id,omitempty" update:""`
	Name                    string            `json:"ad_name,omitempty" update:""`
	IdentityType            string            `json:"identity_type,omitempty" update:""`
	IdentityId              string            `json:"identity_id,omitempty" update:""`
	CatalogId               string            `json:"catalog_id,omitempty"`
	ProductSetId            string            `json:"product_set_id,omitempty"`
	ProductSpecificType     string            `json:"product_specific_type,omitempty"`
	AdFormat                string            `json:"ad_format,omitempty"`
	AdText                  string            `json:"ad_text,omitempty" update:""`
	DynamicFormat           AdDynamicFormat   `json:"dynamic_format,omitempty"`
	VerticalVideoStrategy   string            `json:"vertical_video_strategy,omitempty"`
	DynamicDestination      string            `json:"dynamic_destination,omitempty"`
	CTA                     AdCta             `json:"call_to_action,omitempty"`
	AigcDisclosureType      string            `json:"aigc_disclosure_type,omitempty"`
	ShoppingAdsFallbackType string            `json:"shopping_ads_fallback_type,omitempty"`
	OperationStatus         AdOperationStatus `json:"operation_status,omitempty"`
	SecondaryStatus         string            `json:"secondary_status,omitempty"`

	VideoId                   string   `json:"video_id,omitempty" update:""`
	ImageIds                  []string `json:"image_ids,omitempty" update:""`
	ShoppingAdsVideoPackageId string   `json:"shopping_ads_video_package_id,omitempty"`
	LandingPageUrl            string   `json:"landing_page_url,omitempty"`
}

func (a *Ad) clearForUpdate() {
	fields := reflect.TypeOf(*a)
	values := reflect.ValueOf(a).Elem()

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		if _, ok := field.Tag.Lookup("update"); !ok {
			values.FieldByName(field.Name).Set(reflect.Zero(field.Type))
		}
	}
}

func (a *Ad) GetEnabled() bool {
	return a.OperationStatus == AdOperationStatus_ENABLE
}

func (a *Ad) SetIdentity(id string) {
	a.IdentityId = id
	a.IdentityType = "CUSTOMIZED_USER"
}

func (a *Ad) SetMediaAssets(videoId string, images []string) {
	a.VideoId = videoId
	a.ImageIds = images
	a.AdFormat = "SINGLE_VIDEO"
}

func (a *Ad) SetEnabled(enabled bool) {
	if enabled {
		a.OperationStatus = AdOperationStatus_ENABLE
	} else {
		a.OperationStatus = AdOperationStatus_DISABLE
	}
}

func (a *Ad) SetAiGeneratedContent(flag bool) {
	if flag {
		a.AigcDisclosureType = "SELF_DISCLOSURE"
	} else {
		a.AigcDisclosureType = "NOT_DECLARED"
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

func (a *Ad) WithGeneratedCatalogVideo() {
	a.VerticalVideoStrategy = "CATALOG_VIDEOS"
	a.ShoppingAdsFallbackType = "SHOPPING_ADS"
	a.DynamicDestination = "UNSET"
}

func (a *Ad) WithCustomUrlAndAssets(landingPageUrl, videoId string, imageIds ...string) {
	a.VerticalVideoStrategy = "SINGLE_VIDEO"
	a.ShoppingAdsFallbackType = "CUSTOM"
	a.DynamicDestination = "DLP"

	a.VideoId = videoId
	a.ImageIds = imageIds
	a.LandingPageUrl = landingPageUrl
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
	} else {
		if len(ad.OperationStatus) > 0 {
			if e := updateAdStatus(accountId, ad.Id, ad.OperationStatus); e != nil {
				return nil, e
			}
		}

		ad.clearForUpdate()
		// update other fields
		payload := &createAdRequest{
			AccountId:   accountId,
			AdGroupId:   adGroupId,
			PatchUpdate: new(bool),
			Creatives:   []*Ad{ad},
		}
		*payload.PatchUpdate = true

		req := newPostRequest(
			urlAdUpdate,
			payload,
		)

		if _, err := fetch[emptyResult](req); nil != err {
			return nil, err
		}

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
