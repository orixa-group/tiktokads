package tiktokads

import (
	"testing"
)

func TestGetRecommendedShoppingCtas(t *testing.T) {
	initTestSession()

	ctas, err := GetRecommendedShoppingCtas(getTestAccount())
	assertNotEmptyResult(t, ctas, err)

	for _, cta := range ctas {
		if len(cta.AssetIds) == 0 {
			t.Error("missing cta ids")
		}
		if len(cta.Title) == 0 {
			t.Error("missing cta title")
		}
	}
}
