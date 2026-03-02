package tiktokads

type Identity struct {
	Id              string                  `json:"identity_id"`
	Type            IdentityType            `json:"identity_type"`
	UserName        string                  `json:"identity_username"`
	DisplayName     string                  `json:"display_name"`
	ProfileImage    string                  `json:"profile_image"`
	AvailableStatus IdentityAvailableStatus `json:"available_status"`
}

func (id *Identity) IsAvailable() bool {
	return len(id.AvailableStatus) == 0 || id.AvailableStatus == IdentityAvailableStatus_AVAILABLE
}

func (id *Identity) IsCustomizedUser() bool {
	return id.Type == IdentityType_CUSTOMIZED_USER
}

func GetAccountIdentities(accountId string) ([]*Identity, error) {
	req := newGetRequest(
		urlIdentityGet,
		withAccountId(accountId),
	)

	return fetchAllPages[Identity](req, 100)
}
