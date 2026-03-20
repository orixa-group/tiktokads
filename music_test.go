package tiktokads

import (
	"testing"
)

func TestGetRecommendedCatalogMusics(t *testing.T) {
	initTestSession()
	businessId := "7489850869144485895"
	catalogId := "7595212727484827393"

	musics, err := GetRecommendedCatalogMusics(getTestAccount(), businessId, catalogId)
	assertNotEmptyResult(t, musics, err)
}
