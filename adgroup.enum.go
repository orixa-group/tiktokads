package tiktokads

type AdGroupShoppingAdsType string

const (
	AdGroupShoppingAdsType_PRODUCT_SHOPPING_ADS AdGroupShoppingAdsType = "PRODUCT_SHOPPING_ADS"
	AdGroupShoppingAdsType_VIDEO                AdGroupShoppingAdsType = "VIDEO"
)

type AdGroupProductSource string

const (
	AdGroupProductSource_CATALOG AdGroupProductSource = "CATALOG"
)

type AdGroupGender string

const (
	AdGroupGender_MALE   AdGroupGender = "GENDER_MALE"
	AdGroupGender_FEMALE AdGroupGender = "GENDER_FEMALE"
	AdGroupGender_ALL    AdGroupGender = "GENDER_UNLIMITED"
)

type AdGroupAgeGroup string

const (
	AdGroupAgeGroup_13_17  AdGroupAgeGroup = "AGE_13_17"
	AdGroupAgeGroup_18_24  AdGroupAgeGroup = "AGE_18_24"
	AdGroupAgeGroup_25_34  AdGroupAgeGroup = "AGE_25_34"
	AdGroupAgeGroup_35_44  AdGroupAgeGroup = "AGE_35_44"
	AdGroupAgeGroup_45_54  AdGroupAgeGroup = "AGE_45_54"
	AdGroupAgeGroup_55_100                 = "AGE_55_100"
)

type AdGroupScheduleType string

const (
	AdGroupScheduleType_FROM_NOW AdGroupScheduleType = "SCHEDULE_FROM_NOW"
)

type AdGroupBillingEvent string

const (
	AdGroupBillingEvent_CPC  AdGroupBillingEvent = "CPC"
	AdGroupBillingEvent_OCPM AdGroupBillingEvent = "OCPM"
)

type AdGroupPlacement string

const (
	AdGroupPlacement_TIKTOK AdGroupPlacement = "PLACEMENT_TIKTOK"
)

type AdGroupPromotionType string

const (
	AdGroupPromotionType_WEBSITE AdGroupPromotionType = "WEBSITE"
	AdGroupPromotionType_PRODUCT AdGroupPromotionType = "PSA_PRODUCT"
)

type AdGroupOptimizationGoal string

const (
	AdGroupOptimizationGoal_CONVERT AdGroupOptimizationGoal = "CONVERT"
	AdGroupOptimizationGoal_CLICK   AdGroupOptimizationGoal = "CLICK"
	AdGroupOptimizationGoal_VALUE   AdGroupOptimizationGoal = "VALUE"
)

type AdGroupOptimizationEvent string

const (
	AdGroupOptimizationEvent_BUTTON   AdGroupOptimizationEvent = "BUTTON"
	AdGroupOptimizationEvent_SHOPPING AdGroupOptimizationEvent = "SHOPPING"
)

type AdGroupBidType string

const (
	AdGroupBidType_CUSTOM AdGroupBidType = "BID_TYPE_CUSTOM"
	AdGroupBidType_NO_BID AdGroupBidType = "BID_TYPE_NO_BID"
)

type DeliveryMode string

const (
	DeliveryMode_STANDARD DeliveryMode = "STANDARD"
)
