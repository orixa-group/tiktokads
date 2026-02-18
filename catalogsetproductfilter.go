package tiktokads

import (
	"slices"
)

type ProductFilters struct {
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
	IsIncluded   bool     `json:"is_included"`
	CanUpdate    bool     `json:"can_update"`
}

func newProductFiltersFromApi(conds *condition) *ProductFilters {
	filters := &ProductFilters{
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
				filters.IsIncluded = includeRuleFound[0]
			}
		}
	}

	return filters
}

func newConditionFromProductFilters(filters *ProductFilters) *condition {
	res := &condition{
		And: make([]*rule, 1),
	}

	if nb := len(filters.ProductIds); nb > 0 {
		res.And[0] = newRule("item_group_id", filters.IsIncluded, filters.ProductIds)
	} else if nb = len(filters.ProductTypes); nb > 0 {
		res.And[0] = newRule("product_type", filters.IsIncluded, filters.ProductTypes)
	} else if nb = len(filters.Conditions); nb > 0 {
		res.And[0] = newRule("condition", filters.IsIncluded, filters.Conditions)
	} else if nb = len(filters.Brands); nb > 0 {
		res.And[0] = newRule("brand", filters.IsIncluded, filters.Brands)
	} else if nb = len(filters.Categories); nb > 0 {
		res.And[0] = newRule("category", filters.IsIncluded, filters.Categories)
	} else if nb = len(filters.Labels0); nb > 0 {
		res.And[0] = newRule("custom_label_0", filters.IsIncluded, filters.Labels0)
	} else if nb = len(filters.Labels1); nb > 0 {
		res.And[0] = newRule("custom_label_1", filters.IsIncluded, filters.Labels1)
	} else if nb = len(filters.Labels2); nb > 0 {
		res.And[0] = newRule("custom_label_2", filters.IsIncluded, filters.Labels2)
	} else if nb = len(filters.Labels3); nb > 0 {
		res.And[0] = newRule("custom_label_3", filters.IsIncluded, filters.Labels3)
	} else if nb = len(filters.Labels4); nb > 0 {
		res.And[0] = newRule("custom_label_4", filters.IsIncluded, filters.Labels4)
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
