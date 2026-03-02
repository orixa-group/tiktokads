package tiktokads

// NewCatalogAd Build ad from catalog. Video will be generated from product catalog items
// https://business-api.tiktok.com/portal/docs?id=1750361698613249 "Ad format as Catalog Video"
func NewCatalogAd(name, identityId, catalogId string) *Ad {
	return &Ad{
		Name:                    name,
		IdentityType:            "CUSTOMIZED_USER",
		IdentityId:              identityId,
		CatalogId:               catalogId,
		DynamicFormat:           AdDynamicFormat_UNSET,
		DynamicDestination:      "UNSET",
		AdFormat:                "SINGLE_VIDEO",
		VerticalVideoStrategy:   "CATALOG_VIDEOS",
		CTA:                     AdCta_SHOP_NOW,
		ProductSpecificType:     "ALL",
		AigcDisclosureType:      "NOT_DECLARED",
		ShoppingAdsFallbackType: "SHOPPING_ADS",
	}
}
