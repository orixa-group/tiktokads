package tiktokads

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *rule) UnmarshalJSON(bytes []byte) error {
	data := map[string]any{}

	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	if val, ok := data["field"].(string); ok {
		s.Field = val
	}
	if val, ok := data["operation"].(string); ok {
		s.Operation = val
	}

	if val, ok := data["value"]; ok {
		// Values can either be string or array of string
		// try unmarshall in string first
		buf, _ := json.Marshal(val)
		var singleValue string

		if err := json.Unmarshal(buf, &singleValue); err == nil {
			s.Values = []string{singleValue}
		} else {
			json.Unmarshal(buf, &s.Values)
		}
	}

	return nil
}

type rule struct {
	Field     string   `json:"field"`
	Operation string   `json:"operation"`
	Values    []string `json:"value"`
}

type condition struct {
	And []*rule `json:"and,omitempty"`
	Or  []*rule `json:"or,omitempty"`
}

type catalogSet struct {
	Id           string     `json:"product_set_id,omitempty"`
	Name         string     `json:"product_set_name,omitempty"`
	Conditions   *condition `json:"conditions,omitempty"`
	ProductCount int        `json:"product_count,omitempty"`
	// for update
	BcId      string `json:"bc_id,omitempty"`
	CatalogId string `json:"catalog_id,omitempty"`
}

type ProductSet struct {
	Id           string          `json:"id"`
	Name         string          `json:"name"`
	Filters      *ProductFilters `json:"filters"`
	ProductCount int             `json:"product_count"`
}

func newProductSetFromApi(c *catalogSet) *ProductSet {
	return &ProductSet{
		Id:           c.Id,
		Name:         c.Name,
		ProductCount: c.ProductCount,
		Filters:      newProductFiltersFromApi(c.Conditions),
	}
}

func GetCatalogSets(businessId, catalogId string) ([]*ProductSet, error) {
	req := newGetRequest(
		urlCatalogSetGet,
		withBusinessId(businessId),
		withCatalogId(catalogId),
	)
	sets, err := fetch[listResult[catalogSet]](req)

	if nil != err {
		return nil, err
	}

	results := make([]*ProductSet, len(sets.List))

	for i, set := range sets.List {
		results[i] = newProductSetFromApi(set)
	}

	return results, nil
}

func UpdateCatalogSet(businessId, catalogId string, productSet *ProductSet) (*ProductSet, error) {
	body := &catalogSet{
		Id:         productSet.Id,
		Name:       productSet.Name,
		Conditions: newConditionFromProductFilters(productSet.Filters),
		BcId:       businessId,
		CatalogId:  catalogId,
	}

	type response struct {
		ProductSetId string `json:"product_set_id"`
	}

	var req *http.Request
	if len(productSet.Id) > 0 {
		req = newPostRequest(
			urlCatalogSetUpdate,
			body,
		)
	} else {
		req = newPostRequest(
			urlCatalogSetCreate,
			body,
		)
	}

	resp, err := fetch[response](req)

	if nil != err {
		return nil, err
	} else {
		sets, _ := GetCatalogSets(businessId, catalogId)
		for _, set := range sets {
			if set.Id == resp.ProductSetId {
				return set, nil
			}
		}
	}

	return nil, fmt.Errorf("product set id '%s' not found after update", resp.ProductSetId)
}
