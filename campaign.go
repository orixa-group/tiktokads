package tiktokads

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

type Campaign struct {
	Id                   string                       `json:"campaign_id,omitempty"`
	Name                 string                       `json:"campaign_name,omitempty"`
	Budget               float64                      `json:"budget,omitempty"`
	CampaignType         string                       `json:"campaign_type,omitempty"`
	SalesDestination     CampaignSalesDestination     `json:"sales_destination,omitempty"`
	VirtualObjectiveType CampaignVirtualObjectiveType `json:"virtual_objective_type,omitempty"`
	ProductSource        CampaignProductSource        `json:"campaign_product_source,omitempty"`
	OperationStatus      CampaignOperationStatus      `json:"operation_status,omitempty"`
	ObjectiveType        CampaignObjectiveType        `json:"objective_type,omitempty"`
	BudgetMode           CampaignBudgetMode           `json:"budget_mode,omitempty"`
	BudgetOptimization   bool                         `json:"budget_optimize_on"`
	CatalogEnabled       bool                         `json:"catalog_enabled"`

	// Smart+
	CatalogType string `json:"catalog_type"`
}

type campaignCreate struct {
	AdvertiserId string `json:"advertiser_id,omitempty"`
	RequestId    int64  `json:"request_id,omitempty,string"`
	Campaign
}

type campaignUpdateRequest struct {
	AdvertiserId string  `json:"advertiser_id,omitempty"`
	CampaignId   string  `json:"campaign_id,omitempty"`
	Name         string  `json:"campaign_name,omitempty"`
	Budget       float64 `json:"budget,omitempty"`
}

func newCampaignUpdateRequest(accountId string, c *Campaign) *campaignUpdateRequest {
	return &campaignUpdateRequest{
		AdvertiserId: accountId,
		CampaignId:   c.Id,
		Name:         c.Name,
		Budget:       c.Budget,
	}
}

type campaignCreateResponse struct {
	AdvertiserId string `json:"advertiser_id,omitempty"`
	CampaignId   string `json:"campaign_id,omitempty"`
}

// GetCampaigns Get campaign accounts
func GetCampaigns(accountId string) ([]*Campaign, error) {
	req := newGetRequest(
		urlCampaignsFetch,
		withAccountId(accountId),
	)

	return fetchAllPages[Campaign](req, 1000)
}

// GetCampaign Get campaign from id
func GetCampaign(accountId, campaignId string) (*Campaign, error) {
	req := newGetRequest(
		urlCampaignsFetch,
		withAccountId(accountId),
		withFiltering(map[string]any{
			"campaign_ids": []string{campaignId},
		}),
	)

	if campaigns, err := fetchAllPages[Campaign](req, 1000); err != nil {
		return nil, err
	} else if len(campaigns) == 0 {
		return nil, campaignNotFoundError
	} else if len(campaigns) > 1 {
		return nil, fmt.Errorf("unexpected campaign count: %d", len(campaigns))
	} else {
		return campaigns[0], nil
	}
}

func DeleteCampaign(accountId, campaignId string) error {
	return updateCampaignStatus(accountId, campaignId, CampaignOperationStatus_DELETED)
}

func updateCampaignStatus(accountId, campaignId string, status CampaignOperationStatus) error {
	req := newPostRequest(
		urlCampaignStatusUpdate,
		&map[string]any{
			"advertiser_id":    accountId,
			"campaign_ids":     []string{campaignId},
			"operation_status": status,
		},
	)

	_, err := fetch[emptyResult](req)

	return err
}

// UpdateCampaign
func updateCampaign(accountId string, campaign *Campaign, isSmart bool) (*Campaign, error) {
	var id = campaign.Id

	urlsCreate := map[bool]string{
		false: urlCampaignCreate,
		true:  urlCampaignSmartCreate,
	}
	urlsUpdate := map[bool]string{
		false: urlCampaignUpdate,
		true:  urlCampaignSmartUpdate,
	}

	if len(id) == 0 {
		payload := &campaignCreate{
			AdvertiserId: accountId,
			Campaign:     *campaign,
		}
		if isSmart {
			payload.RequestId = rand.Int64()
		}
		req := newPostRequest(
			urlsCreate[isSmart],
			payload,
		)

		if created, err := fetch[campaignCreateResponse](req); nil != err {
			return nil, err
		} else {
			id = created.CampaignId
		}
	} else {
		req := newPostRequest(
			urlsUpdate[isSmart],
			&campaignCreate{
				AdvertiserId: accountId,
				Campaign:     *campaign,
			},
		)

		if _, err := fetch[emptyResult](req); nil != err {
			return nil, err
		}

		if len(campaign.OperationStatus) > 0 {
			if err := updateCampaignStatus(accountId, campaign.Id, campaign.OperationStatus); nil != err {
				return nil, err
			}
		}
	}

	return waitForCampaign(accountId, id)
}

func UpdateCampaign(accountId string, campaign *Campaign) (*Campaign, error) {
	return updateCampaign(accountId, campaign, false)
}

func UpdateSmartCampaign(accountId string, campaign *Campaign) (*Campaign, error) {
	return updateCampaign(accountId, campaign, true)
}

func waitForCampaign(accountId, campaignId string) (*Campaign, error) {
	for i := float64(1); i <= 10; i++ {
		if campaign, err := GetCampaign(accountId, campaignId); errors.Is(err, campaignNotFoundError) {
			time.Sleep(time.Duration(math.Pow(1.44, i)*1000) * time.Millisecond)
		} else if nil != err {
			return nil, err
		} else {
			return campaign, nil
		}
	}

	return nil, searchTimeoutError
}
