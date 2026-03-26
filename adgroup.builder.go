package tiktokads

import "time"

// https://business-api.tiktok.com/portal/docs?id=1750361698613249#item-link-1.%20Create%20a%20campaign
func NewShoppingAdGroup(name, businessId, pixelId, catalogId string) *AdGroup {
	return &AdGroup{
		BidType: AdGroupBidType_NO_BID,

		BillingEvent: AdGroupBillingEvent_OCPM,
		CatalogId:    catalogId,
		Gender:       AdGroupGender_ALL,
		Name:         name,

		OptimizationGoal:  AdGroupOptimizationGoal_CONVERT,
		OptimizationEvent: AdGroupOptimizationEvent_SHOPPING,
		PromotionType:     "WEBSITE",

		PixelId: pixelId,

		PlacementType: "PLACEMENT_TYPE_NORMAL",
		Placements:    []AdGroupPlacement{AdGroupPlacement_TIKTOK},
		ProductSource: AdGroupProductSource_CATALOG,
		DeliveryMode:  DeliveryMode_STANDARD,
		ScheduleType:  AdGroupScheduleType_FROM_NOW,

		ShoppingAdsType: AdGroupShoppingAdsType_VIDEO,

		StartTime:             time.Now().Format(time.DateTime),
		VideoDownloadDisabled: true,

		BcId: businessId,

		BudgetMode:                 "BUDGET_MODE_DYNAMIC_DAILY_BUDGET",
		ShoppingAdsRetargetingType: "OFF",
	}
}

func NewSmartAdGroup(businessId, catalogId, pixelId string) *AdGroup {
	return &AdGroup{
		BcId:              businessId,
		CatalogId:         catalogId,
		PromotionType:     AdGroupPromotionType_WEBSITE,
		OptimizationGoal:  AdGroupOptimizationGoal_CONVERT,
		OptimizationEvent: AdGroupOptimizationEvent_SHOPPING,
		PixelId:           pixelId,
		BillingEvent:      AdGroupBillingEvent_OCPM,
		ScheduleType:      AdGroupScheduleType_FROM_NOW,

		PlacementType: "PLACEMENT_TYPE_NORMAL",
		Placements:    []AdGroupPlacement{AdGroupPlacement_TIKTOK},

		BidType:                   AdGroupBidType_NO_BID,
		DeepBidType:               "VO_HIGHEST_VALUE",
		TargetingOptimizationMode: "AUTOMATIC",
		StartTime:                 time.Now().Format(time.DateTime),
	}
}
