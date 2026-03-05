package tiktokads

// NewCatalogAd Build ad from catalog. Video will be generated from product catalog items
// https://business-api.tiktok.com/portal/docs?id=1750361698613249 "Ad format as Catalog Video"
func newShoppingAds(name, identityId, catalogId string) *Ad {
	return &Ad{
		Name:                name,
		IdentityType:        "CUSTOMIZED_USER",
		IdentityId:          identityId,
		CatalogId:           catalogId,
		AdFormat:            "SINGLE_VIDEO",
		CTA:                 AdCta_SHOP_NOW,
		ProductSpecificType: "ALL",
		DynamicFormat:       AdDynamicFormat_UNSET,
	}
}

func NewCatalogAd(name, identityId, catalogId string) *Ad {
	ad := newShoppingAds(name, identityId, catalogId)
	ad.WithGeneratedCatalogVideo()
	ad.AigcDisclosureType = "NOT_DECLARED"

	return ad

}

func NewCatalogAdWithCustomVideo(name, identityId, catalogId, lpUrl, videoId string, imageIds ...string) *Ad {
	ad := newShoppingAds(name, identityId, catalogId)
	ad.WithCustomUrlAndAssets(lpUrl, videoId, imageIds...)

	return ad
}
