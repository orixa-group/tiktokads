package tiktokads

import (
	"cmp"
	"net/url"
	"testing"
)

func TestGetAccountIdentities(t *testing.T) {
	initTestSession()

	ids, err := GetAccountIdentities("7101293445955731458")
	assertNotEmptyResult(t, ids, err)
	available := false
	for _, id := range ids {
		available = available || id.IsAvailable()
		if _, err := url.Parse(id.ProfileImage); nil != err {
			t.Error(err)
		}
		if len(cmp.Or(id.UserName, id.DisplayName)) == 0 {
			t.Error("user name & display name are empty")
		}
	}
	if !available {
		t.Error()
	}
}
