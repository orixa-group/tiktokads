package tiktokads

import "time"

func NewShoppingAdGroup(name, pixelId, catalogId string) *AdGroup {
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
		//ShoppingAdsType: AdGroupShoppingAdsType_PRODUCT_SHOPPING_ADS,

		StartTime:             time.Now().Format(time.DateTime),
		VideoDownloadDisabled: true,

		BcId:   "7489850869144485895",
		Pacing: "PACING_MODE_SMOOTH",

		BudgetMode: "BUDGET_MODE_DYNAMIC_DAILY_BUDGET",
		Budget:     56.78,
	}
}
