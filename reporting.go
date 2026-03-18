package tiktokads

import (
	"encoding/json"
	"time"
)

type reportingDimension struct {
	StatTimeDay time.Time `json:"stat_time_day"`
	CampaignId  string    `json:"campaign_id"`
}

func (r *reportingDimension) UnmarshalJSON(bytes []byte) error {
	var data map[string]string

	json.Unmarshal(bytes, &data)

	r.CampaignId = data["campaign_id"]
	if d, ok := data["stat_time_day"]; ok {
		r.StatTimeDay, _ = time.Parse(time.DateTime, d)
	}

	return nil
}

// ReportRow https://business-api.tiktok.com/portal/docs?id=1759239462689793
type ReportRow struct {
	Dimensions reportingDimension `json:"dimensions"`
	Metrics    tiktokMetric       `json:"metrics"`
}

type tiktokMetric struct {
	Impressions int     `json:"impressions,string"`
	Spend       float32 `json:"spend,string"`
	Clicks      float32 `json:"clicks,string"`
	Conversions float32 `json:"conversion,string,omitempty"`
}

func (tm *tiktokMetric) IsEmpty() bool {
	return tm.Impressions == 0 &&
		tm.Spend == 0 &&
		tm.Clicks == 0 &&
		tm.Conversions == 0
}

func GetCampaignReporting(accountId string, from, to time.Time) ([]*ReportRow, error) {
	var rows []*ReportRow

	// can only retrieve 30-days
	for nextStart := from; nextStart.Unix() <= to.Unix(); {
		end := nextStart.AddDate(0, 0, 30)

		if end.After(to) {
			end = to
		}

		req := newGetRequest(
			urlReporting,
			withAccountId(accountId),
			withQueryStringArrayParam("metrics", []string{
				"impressions",
				"spend",
				"clicks",
				"conversion",
			}),
			withQueryStringArrayParam("dimensions", []string{
				"campaign_id",
				"stat_time_day",
			}),
			withQueryString(map[string]string{
				"report_type": "BASIC",
				"data_level":  "AUCTION_CAMPAIGN",
				"start_date":  nextStart.Format(time.DateOnly),
				"end_date":    end.Format(time.DateOnly),
			}),
		)

		if batch, err := fetchAllPages[ReportRow](req, 1000); nil != err {
			return nil, err
		} else {
			for _, row := range batch {
				if !row.Metrics.IsEmpty() {
					rows = append(rows, row)
				}
			}
		}

		// Move to next period
		nextStart = end.AddDate(0, 0, 1)
	}

	return rows, nil
}
