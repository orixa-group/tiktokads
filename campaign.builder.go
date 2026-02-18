package tiktokads

func NewShoppingCatalogCampaign(name string) *Campaign {
	c := &Campaign{
		Name:                 name,
		VirtualObjectiveType: CampaignVirtualObjectiveType_SALES,
		SalesDestination:     CampaignSalesDestination_WEBSITE,
		ProductSource:        CampaignProductSource_CATALOG,
		OperationStatus:      CampaignOperationStatus_PAUSED,
		ObjectiveType:        CampaignObjectiveType_PRODUCT_SALES,
		//ObjectiveType:      CampaignObjectiveType_WEB_CONVERSIONS,
		BudgetMode:         CampaignBudgetMode_DAY,
		BudgetOptimization: true,
		//CatalogEnabled:     true,
	}

	return c
}
