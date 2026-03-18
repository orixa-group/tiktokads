package tiktokads

import (
	"testing"
	"time"
)

func TestGetCampaignReporting(t *testing.T) {
	initTestSession()

	from, _ := time.Parse(time.DateOnly, "2026-01-01")
	to, _ := time.Parse(time.DateOnly, "2026-03-31")

	rows, err := GetCampaignReporting(getTestAccount(), from, to)
	if nil != err {
		t.Fatal(err)
	}

	assertNotEmptyResult(t, rows, err)
	for _, row := range rows {
		if len(row.Dimensions.CampaignId) == 0 {
			t.Error("dimensions.CampaignId is empty")
		}
		if row.Dimensions.StatTimeDay.Before(from) || row.Dimensions.StatTimeDay.After(to) {
			t.Error("dimensions.StatTimeDay is outside of dateRange")
		}
		if row.Metrics.IsEmpty() {
			t.Error("dimensions.Metrics is empty")
		}
	}
}
