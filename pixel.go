package tiktokads

type Pixel struct {
	Id   string `json:"pixel_id"`
	Code string `json:"pixel_code"`
	Name string `json:"pixel_name"`
}

func GetBusinessPixels(businessId string) ([]*Pixel, error) {
	req := newGetRequest(
		urlBusinessPixelGet,
		withBusinessId(businessId),
	)

	res, err := fetchAllPages[Pixel](req, 100)

	return res, err
}
