package tiktokads

type AdGroup struct {
	Id                         string                   `json:"adgroup_id,omitempty"`
	Name                       string                   `json:"adgroup_name,omitempty"`
	ShoppingAdsType            AdGroupShoppingAdsType   `json:"shopping_ads_type,omitempty"`
	ProductSource              AdGroupProductSource     `json:"product_source,omitempty"`
	OptimizationGoal           AdGroupOptimizationGoal  `json:"optimization_goal,omitempty"`
	OptimizationEvent          AdGroupOptimizationEvent `json:"optimization_event,omitempty"`
	BidStrategy                string                   `json:"bid_strategy,omitempty"`
	BudgetMode                 string                   `json:"budget_mode"`
	Budget                     float32                  `json:"budget"`
	CatalogId                  string                   `json:"catalog_id,omitempty"`
	LocationIds                []string                 `json:"location_ids,omitempty"`
	Languages                  []string                 `json:"languages,omitempty"`
	Gender                     AdGroupGender            `json:"gender,omitempty"`
	AgeGroups                  []AdGroupAgeGroup        `json:"age_groups,omitempty"`
	AudienceIds                []AdGroupAgeGroup        `json:"audience_ids,omitempty"`
	ExcludedAudienceIds        []AdGroupAgeGroup        `json:"excluded_audience_ids,omitempty"`
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
	ShoppingAdsRetargetingType string                   `json:"shopping_ads_retargeting_type"`

	Pacing string `json:"pacing,omitempty"`
}

type adGroupUpdateRequest struct {
	AdvertiserId string `json:"advertiser_id"`
	CampaignId   string `json:"campaign_id,omitempty"` // required only for creation
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

func UpdateAdGroup(accountId, campaignId string, ag *AdGroup) (*AdGroup, error) {
	payload := &adGroupUpdateRequest{
		AdvertiserId: accountId,
		AdGroup:      ag,
	}
	url := urlAdGroupUpdate
	if len(ag.Id) == 0 {
		payload.CampaignId = campaignId
		url = urlAdGroupCreate
	}

	req := newPostRequest(
		url,
		payload,
	)

	updated, err := fetch[AdGroup](req)

	return updated, err
}
