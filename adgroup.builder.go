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
