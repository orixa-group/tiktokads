package tiktokads

type AdDynamicFormat string

const (
	AdDynamicFormat_DYNAMIC_CREATIVE AdDynamicFormat = "DYNAMIC_CREATIVE" // for custom videos
	AdDynamicFormat_UNSET            AdDynamicFormat = "UNSET"            // Catalog-based videos
)

// AdCta https://business-api.tiktok.com/gateway/docs/index?identify_key=c0138ffadd90a955c1f0670a56fe348d1d40680b3c89461e09f78ed26785164b&language=ENGLISH&doc_id=1737174886619138#item-link-Call-to-action%20(call_to_action)
type AdCta string

const (
	AdCta_SHOP_NOW AdCta = "SHOP_NOW"
)

type AdOperationStatus string

const (
	AdOperationStatus_ENABLE  AdOperationStatus = "ENABLE"
	AdOperationStatus_DISABLE AdOperationStatus = "DISABLE"
	AdOperationStatus_DELETE  AdOperationStatus = "DELETE"
)
