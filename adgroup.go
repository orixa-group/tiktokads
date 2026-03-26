package tiktokads

import (
	"cmp"
	"math/rand/v2"
)

type AdGroup struct {
	Id                         string                   `json:"adgroup_id,omitempty"`
	Name                       string                   `json:"adgroup_name,omitempty"`
	ShoppingAdsType            AdGroupShoppingAdsType   `json:"shopping_ads_type,omitempty"`
	ProductSource              AdGroupProductSource     `json:"product_source,omitempty"`
	OptimizationGoal           AdGroupOptimizationGoal  `json:"optimization_goal,omitempty"`
	OptimizationEvent          AdGroupOptimizationEvent `json:"optimization_event,omitempty"`
	BidStrategy                string                   `json:"bid_strategy,omitempty"`
	BudgetMode                 string                   `json:"budget_mode,omitempty"`
	Budget                     float32                  `json:"budget,omitempty"`
	CatalogId                  string                   `json:"catalog_id,omitempty"`
	LocationIds                []string                 `json:"location_ids,omitempty"`
	Languages                  []string                 `json:"languages,omitempty"`
	Gender                     AdGroupGender            `json:"gender,omitempty"`
	AgeGroups                  []AdGroupAgeGroup        `json:"age_groups,omitempty"`
	AudienceIds                []string                 `json:"audience_ids,omitempty"`
	ExcludedAudienceIds        []string                 `json:"excluded_audience_ids,omitempty"`
	ScheduleType               AdGroupScheduleType      `json:"schedule_type,omitempty"`
	StartTime                  string                   `json:"schedule_start_time,omitempty"`
	BillingEvent               AdGroupBillingEvent      `json:"billing_event,omitempty"`
	PlacementType              string                   `json:"placement_type,omitempty"`
	Placements                 []AdGroupPlacement       `json:"placements,omitempty"`
	PromotionType              AdGroupPromotionType     `json:"promotion_type,omitempty"`
	BidType                    AdGroupBidType           `json:"bid_type,omitempty"`
	DeliveryMode               DeliveryMode             `json:"delivery_mode,omitempty"`
	BcId                       string                   `json:"catalog_authorized_bc_id,omitempty"`
	VideoDownloadDisabled      bool                     `json:"video_download_disabled,omitempty"`
	PixelId                    string                   `json:"pixel_id,omitempty"`
	ShoppingAdsRetargetingType string                   `json:"shopping_ads_retargeting_type,omitempty"`

	Pacing string `json:"pacing,omitempty"`

	// Smart+ specific fields
	TargetingSpec             *adGroupTargetingSpec `json:"targeting_spec,omitempty"`
	DeepBidType               string                `json:"deep_bid_type,omitempty"`
	TargetingOptimizationMode string                `json:"targeting_optimization_mode,omitempty"`
}

type adGroupTargetingSpec struct {
	LocationIds []string `json:"location_ids,omitempty"`
}

func (a *AdGroup) SetLocationIds(locationIds []string, isSmart bool) {
	if isSmart {
		a.TargetingSpec = cmp.Or(a.TargetingSpec, &adGroupTargetingSpec{})
		a.TargetingSpec.LocationIds = locationIds
	} else {
		a.LocationIds = locationIds
	}
}

type adGroupUpdateRequest struct {
	AdvertiserId string `json:"advertiser_id"`
	CampaignId   string `json:"campaign_id,omitempty"`       // required only for creation
	RequestId    int64  `json:"request_id,omitempty,string"` // required only for smart+
	*AdGroup
}

func GetAdGroups(accountId, campaignId string) ([]*AdGroup, error) {
	req := newGetRequest(
		urlAdGroupGet,
		withAccountId(accountId),
		withFiltering(map[string]any{
			"campaign_ids": []string{campaignId},
		}),
	)

	return fetchAllPages[AdGroup](req, 1000)
}

func updateAdgroup(accountId, campaignId string, ag *AdGroup, isSmart bool) (*AdGroup, error) {
	urlCreate := map[bool]string{
		false: urlAdGroupCreate,
		true:  urlSmartAdGroupCreate,
	}
	urlUpdate := map[bool]string{
		false: urlAdGroupUpdate,
		true:  urlSmartAdGroupUpdate,
	}

	payload := &adGroupUpdateRequest{
		AdvertiserId: accountId,
		AdGroup:      ag,
	}
	if isSmart {
		payload.RequestId = rand.Int64()
	}
	url := urlUpdate[isSmart]
	if len(ag.Id) == 0 {
		payload.CampaignId = campaignId
		url = urlCreate[isSmart]
	}

	req := newPostRequest(
		url,
		payload,
	)

	updated, err := fetch[AdGroup](req)

	return updated, err
}

func UpdateAdGroup(accountId, campaignId string, ag *AdGroup) (*AdGroup, error) {
	return updateAdgroup(accountId, campaignId, ag, false)
}

func UpdateSmartAdgroup(accountId, campaignId string, ad *AdGroup) (*AdGroup, error) {
	return updateAdgroup(accountId, campaignId, ad, true)
}
