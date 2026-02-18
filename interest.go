package tiktokads

import (
	"strings"
)

type Interest struct {
	Id   string `json:"interest_category_id"`
	Name string `json:"interest_category_name"`
}

func GetInterests(accountId, query, language string) ([]*Interest, error) {
	keywords := strings.Fields(query)
	//keywords = []string{query}
	subTypes := []string{"GENERAL_INTEREST"}

	if len(keywords) > 1 {
		subTypes[0] = "ADDITIONAL_INTEREST"
	}

	req := newGetRequest(
		urlInterestsGet,
		withAccountId(accountId),
		withQueryStringArrayParam("keywords", keywords),
		//withQueryStringArrayParam("sub_targeting_types", subTypes),
		withQueryString(map[string]string{
			"mode": "SEMANTIC_RECOMMEND",
			//"mode":     "FUZZ_MATCH",
			"language": strings.ToLower(language),
			//"keyword":  query,
			"limit": "50",
		}),
	)

	return fetchAllPages[Interest](req, 100)
}
