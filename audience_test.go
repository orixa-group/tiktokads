package tiktokads

import (
	"testing"
)

func TestGetAudiences(t *testing.T) {
	initTestSession()

	audiences, err := GetAudiences(getTestAccount())
	assertNotEmptyResult(t, audiences, err)
}
