package tiktokads

type BusinessCenter struct {
	Id       string `json:"bc_id"`
	Name     string `json:"name"`
	Currency string `json:"currency"`
}

type businessCenterApi[L BusinessCenter] struct {
	BcInfo *L `json:"bc_info"`
}

func GetBusinesses() ([]*BusinessCenter, error) {
	req := newGetRequest(urlBusinessCentersFetch)

	res, err := fetchAllPages[businessCenterApi[BusinessCenter]](req, 50)

	var result []*BusinessCenter

	for _, row := range res {
		result = append(result, row.BcInfo)
	}

	return result, err
}
