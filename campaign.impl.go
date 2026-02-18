package tiktokads

type BaseCampaign interface {
	GetId() string
	SetId(string)
	GetName() string
	SetName(string)
	GetBudget() int
	SetBudget(int)
	GetEnabled() bool
	SetEnabled(bool)
}

func (c *Campaign) GetId() string {
	return c.Id
}

func (c *Campaign) SetId(id string) {
	c.Id = id
}

func (c *Campaign) GetName() string {
	return c.Name
}

func (c *Campaign) SetName(name string) {
	c.Name = name
}

func (c *Campaign) GetBudget() int {
	return int(c.Budget * 100)
}

func (c *Campaign) SetBudget(budget int) {
	c.Budget = float64(budget) / 100
}

func (c *Campaign) GetEnabled() bool {
	return c.OperationStatus == CampaignOperationStatus_ENABLED
}

func (c *Campaign) SetEnabled(enabled bool) {
	if enabled {
		c.OperationStatus = CampaignOperationStatus_ENABLED
	} else {
		c.OperationStatus = CampaignOperationStatus_PAUSED
	}
}

type TikTokCampaign interface {
	SetDynamicBudget(bool)
}

func (c *Campaign) SetDynamicBudget(enabled bool) {
	if enabled {
		c.BudgetMode = CampaignBudgetMode_DYNAMIC_DAILY
	} else {
		c.BudgetMode = CampaignBudgetMode_DAY
	}
}
