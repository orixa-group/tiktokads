package tiktokads

import (
	"cmp"
	"os"
	"strconv"
	"testing"
)

func getTestAccount() string {
	return cmp.Or(os.Getenv("TIKTOKADS_TEST_ACCOUNT"), "7101293445955731458")
}

func TestGetCampaigns(t *testing.T) {
	initTestSession()

	campaigns, err := GetCampaigns(getTestAccount())
	assertNotEmptyResult(t, campaigns, err)

	for _, campaign := range campaigns {
		if id, _ := strconv.ParseInt(campaign.Id, 10, 64); 0 == id {
			t.Error("campaign id should be numeric")
		}
		if campaign.Budget <= 0 {
			t.Error("campaign budget should be positive")
		}
	}
}
