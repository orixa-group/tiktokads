package tiktokads

func (c *CatalogProductStatus) GetProductId() string {
	return c.ProductId
}

func (c *CatalogProductStatus) IsBlocking() bool {
	return c.IssueSeverity == "ERROR"
}

func (c *CatalogProductStatus) GetDescription() string {
	return c.IssueTitle
}

func (c *CatalogProductStatus) GetAttribute() string {
	return ""
}

func (c *CatalogProductStatus) NeedUserAction() bool {
	return c.IsBlocking()
}
