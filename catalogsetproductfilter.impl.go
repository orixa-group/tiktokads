package tiktokads

type ProductFilters interface {
	GetProductIds() []string
	GetProductTypes() []string
	GetConditions() []string
	GetCategories() []string
	GetBrands() []string
	GetLabels0() []string
	GetLabels1() []string
	GetLabels2() []string
	GetLabels3() []string
	GetLabels4() []string
	IsIncluded() bool
}

func (pf *productFilters) GetProductIds() []string {
	return pf.ProductIds
}

func (pf *productFilters) GetProductTypes() []string {
	return pf.ProductTypes
}

func (pf *productFilters) GetConditions() []string {
	return pf.Conditions
}

func (pf *productFilters) GetCategories() []string {
	return pf.Categories
}

func (pf *productFilters) GetBrands() []string {
	return pf.Brands
}

func (pf *productFilters) GetLabels0() []string {
	return pf.Labels0
}

func (pf *productFilters) GetLabels1() []string {
	return pf.Labels1
}

func (pf *productFilters) GetLabels2() []string {
	return pf.Labels2
}

func (pf *productFilters) GetLabels3() []string {
	return pf.Labels3
}

func (pf *productFilters) GetLabels4() []string {
	return pf.Labels4
}

func (pf *productFilters) IsIncluded() bool {
	return pf.Included
}
