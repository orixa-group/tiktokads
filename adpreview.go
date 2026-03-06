package tiktokads

type adPreview struct {
	*Ad
	AdvertiserId  string `json:"advertiser_id,omitempty"`
	PreviewType   string `json:"preview_type,omitempty"`
	ObjectiveType string `json:"objective_type,omitempty"`
	Placement     string `json:"placement,omitempty"`
	PreviewFormat string `json:"preview_format,omitempty"`
}

type AdPreview struct {
	PreviewLink string `json:"preview_link"`
	Iframe      string `json:"iframe"`
}

func PreviewAd(accountId string, ad *Ad) (*AdPreview, error) {
	req := newPostRequest(urlAdPreview, NewAdShoppingPreview(accountId, ad))

	return fetch[AdPreview](req)
}
