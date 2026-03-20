package tiktokads

type Music struct {
	Id       string `json:"music_id"`
	Author   string `json:"author"`
	CoverUrl string `json:"cover_url"`
	Duration int    `json:"duration"`
	Style    string `json:"style"`
	Url      string `json:"url"`
}

func GetRecommendedCatalogMusics(accountId, businessId, catalogId string) ([]*Music, error) {
	req := newGetRequest(
		urlGetMusics,
		withAccountId(accountId),
		withQueryString(map[string]string{
			"music_scene": "CATALOG_CAROUSEL",
			"search_type": "SEARCH_BY_RECOMMEND",
		}),
		withFiltering(map[string]any{
			"catalog_id":               catalogId,
			"catalog_authorized_bc_id": businessId,
		}),
	)

	return fetchAllPages[Music](req, 1000)
}
