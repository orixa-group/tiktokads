package tiktokads

type Audience struct {
	Id         string  `json:"audience_id,omitempty"`
	Name       string  `json:"name,omitempty"`
	Coverage   float64 `json:"cover_num,omitempty"`
	IsValid    bool    `json:"is_valid,omitempty"`
	IsExpiring bool    `json:"is_expiring,omitempty"`
}

func GetAudiences(accountId string) ([]*Audience, error) {
	req := newGetRequest(
		urlAudienceGet,
		withAccountId(accountId),
	)

	return fetchAllPages[Audience](req, 100)
}
