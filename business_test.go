package tiktokads

import (
	"log"
	"strconv"
	"testing"
)

func TestGetBusinesses(t *testing.T) {
	initTestSession()

	businesses, err := GetBusinesses()

	assertNotEmptyResult(t, businesses, err)
	log.Println(LastResponse)

	for _, business := range businesses {
		if id, _ := strconv.ParseInt(business.Id, 10, 64); 0 == id {
			t.Error("business id should be numeric string")
		}
		t.Log(business)
	}
}
