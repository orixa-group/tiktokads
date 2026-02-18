package tiktokads

type Account struct {
	Id   string `json:"advertiser_id"`
	Name string `json:"advertiser_name"`
}

func GetAdAccounts() ([]*Account, error) {
	req := newGetRequest(
		urlAccountsFetch,
		withAppIdAndSecret,
	)

	res, err := fetch[listResult[Account]](req)

	if err != nil {
		return nil, err
	}

	return res.List, nil
}
