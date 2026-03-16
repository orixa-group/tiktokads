package tiktokads

import (
	"slices"
)

type productFilters struct {
	ProductIds   []string `json:"product_ids,omitempty"`
	ProductTypes []string `json:"product_types,omitempty"`
	Conditions   []string `json:"conditions,omitempty"`
	Categories   []string `json:"categories,omitempty"`
	Brands       []string `json:"brands,omitempty"`
	Labels0      []string `json:"labels_0,omitempty" api:"custom_label_0"`
	Labels1      []string `json:"labels_1,omitempty" api:"custom_label_1"`
	Labels2      []string `json:"labels_2,omitempty" api:"custom_label_2"`
	Labels3      []string `json:"labels_3,omitempty" api:"custom_label_3"`
	Labels4      []string `json:"labels_4,omitempty" api:"custom_label_4"`
	Included     bool     `json:"is_included"`
	CanUpdate    bool     `json:"can_update"`
}

func newProductFiltersFromApi(conds *condition) *productFilters {
	filters := &productFilters{
		CanUpdate: true,
	}

	includeOperators := []string{
		"EQUAL",
		"CONTAIN",
		"I_CONTAIN",
	}

	includeRuleFound := []bool{}

	if nil != conds {
		if len(conds.Or) > 0 {
			filters.CanUpdate = false
		} else {
			for _, cond := range conds.And {
				if includeFlag := slices.Contains(includeOperators, cond.Operation); !slices.Contains(includeRuleFound, includeFlag) {
					includeRuleFound = append(includeRuleFound, includeFlag)
				}

				switch cond.Field {
				case "sku_id":
				case "item_group_id":
					filters.ProductIds = append(filters.ProductIds, cond.Values...)
					break
				case "condition":
					filters.Conditions = append(filters.Conditions, cond.Values...)
					break
				case "brand":
					filters.Brands = append(filters.Brands, cond.Values...)
					break
				case "category":
				case "google_product_category":
					filters.Categories = append(filters.Categories, cond.Values...)
					break
				case "product_type":
					filters.ProductTypes = append(filters.ProductTypes, cond.Values...)
					break
				case "custom_label_0":
					filters.Labels0 = append(filters.Labels0, cond.Values...)
					break
				case "custom_label_1":
					filters.Labels1 = append(filters.Labels1, cond.Values...)
					break
				case "custom_label_2":
					filters.Labels2 = append(filters.Labels2, cond.Values...)
					break
				case "custom_label_3":
					filters.Labels3 = append(filters.Labels3, cond.Values...)
					break
				case "custom_label_4":
					filters.Labels4 = append(filters.Labels4, cond.Values...)
					break
				default:
					// non-handled rule, skip edition
					filters.CanUpdate = false
				}
			}

			// Do not support heterogeneous include/exclude rules
			if nb := len(includeRuleFound); nb > 1 {
				filters.CanUpdate = false
			} else if nb > 0 {
				filters.Included = includeRuleFound[0]
			}
		}
	}

	return filters
}

func newConditionFromProductFilters(filters ProductFilters) *condition {
	res := &condition{
		And: make([]*rule, 1),
	}

	if nil != filters {
		if nb := len(filters.GetProductIds()); nb > 0 {
			res.And[0] = newRule("item_group_id", filters.IsIncluded(), filters.GetProductIds())
		} else if nb = len(filters.GetProductTypes()); nb > 0 {
			res.And[0] = newRule("product_type", filters.IsIncluded(), filters.GetProductTypes())
		} else if nb = len(filters.GetConditions()); nb > 0 {
			res.And[0] = newRule("condition", filters.IsIncluded(), filters.GetConditions())
		} else if nb = len(filters.GetBrands()); nb > 0 {
			res.And[0] = newRule("brand", filters.IsIncluded(), filters.GetBrands())
		} else if nb = len(filters.GetCategories()); nb > 0 {
			res.And[0] = newRule("category", filters.IsIncluded(), filters.GetCategories())
		} else if nb = len(filters.GetLabels0()); nb > 0 {
			res.And[0] = newRule("custom_label_0", filters.IsIncluded(), filters.GetLabels0())
		} else if nb = len(filters.GetLabels1()); nb > 0 {
			res.And[0] = newRule("custom_label_1", filters.IsIncluded(), filters.GetLabels1())
		} else if nb = len(filters.GetLabels2()); nb > 0 {
			res.And[0] = newRule("custom_label_2", filters.IsIncluded(), filters.GetLabels2())
		} else if nb = len(filters.GetLabels3()); nb > 0 {
			res.And[0] = newRule("custom_label_3", filters.IsIncluded(), filters.GetLabels3())
		} else if nb = len(filters.GetLabels4()); nb > 0 {
			res.And[0] = newRule("custom_label_4", filters.IsIncluded(), filters.GetLabels4())
		}
	}

	return res
}

func newRule(field string, isIncluded bool, values []string) *rule {
	// get operator value depending on include rule (true/false) and exact match(true/false)
	operatorValues := map[bool]map[bool]string{
		true: {
			true:  "EQUAL",
			false: "CONTAIN",
		},
		false: {
			true:  "NOT_EQUAL",
			false: "EXCLUDE",
		},
	}

	return &rule{
		Field:     field,
		Values:    values,
		Operation: operatorValues[isIncluded][false],
	}
}
