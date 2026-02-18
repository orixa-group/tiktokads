package tiktokads

import (
	"testing"
)

func TestGetAudiences(t *testing.T) {
	initTestSession()

	audiences, err := GetAudiences(testAccount)
	assertNotEmptyResult(t, audiences, err)
}
