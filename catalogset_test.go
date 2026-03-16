package tiktokads

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestUnmarshallStringOrArrayOfStringRule(t *testing.T) {
	testCases := map[string][]string{
		`{"and":[{"field":"sku_id","operation":"CONTAIN","value":["2529","2536"]}]}`:  {"2529", "2536"},
		`{"and":[{"field":"availability","operation":"CONTAIN","value":"IN_STOCK"}]}`: {"IN_STOCK"},
	}

	for str, expectedValues := range testCases {
		buf := []byte(str)
		var cond condition
		if e := json.Unmarshal(buf, &cond); e != nil {
			t.Fatal(e)
		} else if len(cond.And) < 1 {
			t.Fatal("expected 1 and-condition")
		} else if !reflect.DeepEqual(expectedValues, cond.And[0].Values) {
			t.Fatal("rule values not equals")
		}
	}
}

func TestGetCatalogSets(t *testing.T) {
	initTestSession()
	businessId := "7489850869144485895"
	catalogId := "7594833492484146945"
	sets, err := GetCatalogSets(businessId, catalogId)
	assertNotEmptyResult(t, sets, err)

	for _, set := range sets {
		if set.ProductCount == 0 {
			t.Error("product count is zero")
		}
	}
}

// TestUpdateCatalogSet
func TestUpdateCatalogSet(t *testing.T) {
	initTestSession()
	businessId := "7489850869144485895"
	catalogId := "7595212727484827393"

	productSet := &ProductSet{
		Name: "[TestUnit] To delete",
		Filters: &productFilters{
			ProductIds: []string{"123", "456"},
			Included:   true,
		},
	}
	createdSet, err := UpdateCatalogSet(businessId, catalogId, productSet)

	if err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(createdSet.Filters, productSet.Filters) {
		t.Error("filters not equals after create")
	}

	createdSet.Name = fmt.Sprintf("[TestUnit] updated %s", time.Now().Format(time.DateTime))
	createdSet.Filters.ProductIds = append(createdSet.Filters.ProductIds, "789")

	updatedSet, err := UpdateCatalogSet(businessId, catalogId, createdSet)
	if err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(updatedSet.Filters, createdSet.Filters) {
		t.Error("filters not equals after update")
	}
}
