package tiktokads

type CampaignVirtualObjectiveType string

const (
	CampaignVirtualObjectiveType_SALES CampaignVirtualObjectiveType = "SALES"
)

type CampaignSalesDestination string

const (
	CampaignSalesDestination_SHOP    CampaignSalesDestination = "TIKTOK_SHOP"
	CampaignSalesDestination_WEBSITE CampaignSalesDestination = "WEBSITE"
)

type CampaignOperationStatus string

const (
	CampaignOperationStatus_ENABLED CampaignOperationStatus = "ENABLE"
	CampaignOperationStatus_PAUSED                          = "DISABLE"
	CampaignOperationStatus_DELETED                         = "DELETE"
)

type CampaignBudgetMode string

const (
	CampaignBudgetMode_DAY           CampaignBudgetMode = "BUDGET_MODE_DAY"
	CampaignBudgetMode_DYNAMIC_DAILY                    = "BUDGET_MODE_DYNAMIC_DAILY_BUDGET"
)

type CampaignProductSource string

const (
	CampaignProductSource_CATALOG CampaignProductSource = "CATALOG"
	CampaignProductSource_STORE                         = "STORE"
)

type CampaignObjectiveType string

const (
	CampaignObjectiveType_PRODUCT_SALES   CampaignObjectiveType = "PRODUCT_SALES"
	CampaignObjectiveType_WEB_CONVERSIONS CampaignObjectiveType = "WEB_CONVERSIONS"
)
