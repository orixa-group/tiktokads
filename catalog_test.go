package tiktokads

import (
	"log"
	"strconv"
	"testing"
)

func TestCreateCatalog(t *testing.T) {
	initTestSession()

	businessId := "7489850869144485895"

	c := NewShoppingCatalog(businessId, "[TestUnit] Demo", "EUR", "FR")

	created, err := CreateCatalog(c)

	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(created)
	}
}

func TestGetCatalogs(t *testing.T) {
	initTestSession()
	businessId := "7489850869144485895"

	catalogs, err := GetCatalogs(businessId)
	assertNotEmptyResult(t, catalogs, err)
	for _, catalog := range catalogs {
		if id, _ := strconv.ParseInt(catalog.Id, 10, 64); id == 0 {
			t.Error("catalog id should be numeric")
		}
		if nil == catalog.Business {
			t.Error("missing business informations")
		}
		if false == catalog.IsAdCreationEligible() {
			t.Error("ad creation not eligible ")
		}
	}
}

func TestDeleteCatalog(t *testing.T) {
	initTestSession()
	businessId := "7489850869144485895"
	c := NewShoppingCatalog(businessId, "[TestUnit] To Delete", "EUR", "FR")
	created, err := CreateCatalog(c)
	if nil != err {
		t.Fatal(err)
	}

	if e := DeleteCatalog(businessId, created.Id); nil != e {
		t.Fatal(e)
	}
}

func TestCatalogAddPixel(t *testing.T) {
	initTestSession()
	businessId := "7489850869144485895"
	catalogId := "7595212727484827393"
	//catalogId := "7594833492484146945"
	accountId := "7101293445955731458"
	//pixelCode := "D5JR9MBC77UDMGSG1UBG"
	pixelCode := "D5JR9MBC77UDMGSG1UBG"

	if e := CatalogAddPixel(businessId, catalogId, accountId, pixelCode); nil != e {
		log.Println(LastResponse)
		t.Fatal(e)
	}
}
