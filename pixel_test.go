package tiktokads

import (
	"strconv"
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

func TestGetAccountPixels(t *testing.T) {
	initTestSession()

	pixels, err := GetAccountPixels(getTestAccount())
	assertNotEmptyResult(t, pixels, err)
	for _, p := range pixels {
		if _, e := strconv.Atoi(p.Id); nil != e {
			t.Error("pixel id should be numeric")
		}
		if len(p.Code) == 0 {
			t.Error("pixel code is empty")
		}
		if len(p.Name) == 0 {
			t.Error("pixel name is empty")
		}
	}
}
