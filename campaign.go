package tiktokads

import (
	"errors"
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
}

type campaignCreate struct {
	AdvertiserId string `json:"advertiser_id,omitempty"`
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
	if campaigns, err := GetCampaigns(accountId); nil != err {
		return nil, err
	} else {
		for _, campaign := range campaigns {
			if campaign.Id == campaignId {
				return campaign, nil
			}
		}
	}

	return nil, errors.New("campaign-not-found")
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
func UpdateCampaign(accountId string, campaign *Campaign) (*Campaign, error) {
	var id = campaign.Id
	if len(id) == 0 {
		req := newPostRequest(
			urlCampaignCreate, &campaignCreate{
				AdvertiserId: accountId,
				Campaign:     *campaign,
			},
		)

		if created, err := fetch[campaignCreateResponse](req); nil != err {
			return nil, err
		} else {
			time.Sleep(5 * time.Second)
			id = created.CampaignId
		}
	} else {
		req := newPostRequest(
			urlCampaignUpdate,
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
		time.Sleep(5 * time.Second)
	}

	return GetCampaign(accountId, id)
}
