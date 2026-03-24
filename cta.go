package tiktokads

type DynamicCta struct {
	AssetIds []string `json:"asset_ids"`
	Title    string   `json:"asset_content"`
}

// GetRecommendedShoppingCtas CTA for Upgraded Smart+ Ads & shopping context
// https://business-api.tiktok.com/portal/docs?id=1739362202742785
func GetRecommendedShoppingCtas(accountId string) ([]*DynamicCta, error) {
	req := newGetRequest(
		urlGetCtas,
		withAccountId(accountId),
		withQueryString(map[string]string{
			"content_type":   "LANDING_PAGE",
			"asset_type":     "CTA_NORMAL",
			"objective_type": "PRODUCT_SALES",
		}),
	)

	return fetchAllPages[DynamicCta](req, 100)
}
