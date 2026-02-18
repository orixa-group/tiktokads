package tiktokads

import (
	"strconv"
	"testing"
)

func TestGetCampaignAdGroups(t *testing.T) {
	initTestSession()

	adGroups, err := GetAdGroups(getTestAccount(), "1857382026716449")
	assertNotEmptyResult(t, adGroups, err)

	for _, adGroup := range adGroups {
		if id, _ := strconv.Atoi(adGroup.Id); 0 == id {
			t.Error("ad group id is empty")
		}

		if len(adGroup.Name) == 0 {
			t.Error("ad group name is empty")
		}
		if 0 == len(adGroup.BidStrategy) {
			t.Error("ad group bid strategy is empty")
		}
		if 0 == len(adGroup.OptimizationGoal) {
			t.Error("ad group optimization goal is empty")
		}
		if 0 == len(adGroup.PromotionType) {
			t.Error("ad group promotion type is empty")
		}
	}
}
