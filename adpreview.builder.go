package tiktokads

func NewAdShoppingPreview(accountId string, ad *Ad) *adPreview {
	if len(ad.Id) > 0 {
		return &adPreview{
			AdvertiserId: accountId,
			Ad: &Ad{
				Id: ad.Id,
			},
			PreviewType: "AD",
		}
	}

	ap := &adPreview{
		AdvertiserId:  accountId,
		Ad:            ad,
		PreviewType:   "ADS_CREATION",
		ObjectiveType: "PRODUCT_SALES",
		Placement:     "PLACEMENT_TIKTOK",

		PreviewFormat: "IN_FEED",
	}

	return ap
}
