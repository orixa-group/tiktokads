package tiktokads

type AdDynamicFormat string

const (
	AdDynamicFormat_DYNAMIC_CREATIVE AdDynamicFormat = "DYNAMIC_CREATIVE" // for custom videos
	AdDynamicFormat_UNSET            AdDynamicFormat = "UNSET"            // Catalog-based videos
)

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
