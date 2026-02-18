package tiktokads

import (
	"log"
	"testing"
)

func TestGetInterests(t *testing.T) {
	initTestSession()

	interests, err := GetInterests(getTestAccount(), "Football français", "FR")
	log.Println(LastResponse)
	assertNotEmptyResult(t, interests, err)
}
