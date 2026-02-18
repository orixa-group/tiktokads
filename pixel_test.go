package tiktokads

import (
	"testing"
)

func TestGetBusinessPixels(t *testing.T) {
	initTestSession()
	businessId := "7489850869144485895"

	pixels, err := GetBusinessPixels(businessId)
	assertNotEmptyResult(t, pixels, err)
	for _, p := range pixels {
		if len(p.Code) == 0 {
			t.Error("pixel code is empty")
		}
		if len(p.Name) == 0 {
			t.Error("pixel name is empty")
		}
	}
}
