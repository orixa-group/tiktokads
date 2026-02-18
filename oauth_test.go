package tiktokads

import (
	"os"
	"testing"
)

const redirectUrl = "https://app.feedcast.ai/oauth/callback"

func initTestSession() {
	SetApiKeys(os.Getenv("TIKTOKADS_API_KEY"), os.Getenv("TIKTOKADS_API_SECRET"))
	SetAccessToken(os.Getenv("TIKTOKADS_ACCESS_TOKEN"))
}

func assertNotEmptyResult[T any](t *testing.T, result []*T, err error) {
	if nil != err {
		t.Fatal(err)
	} else if len(result) == 0 {
		t.Fatal("array result should not be empty")
	}
}

func TestGetOAuthRedirect(t *testing.T) {
	initTestSession()

	u := GetOAuthRedirect(redirectUrl)

	t.Logf("Redirect url: %s", u)
}

func TestGetAccessTokenFromAuthCode(t *testing.T) {
	initTestSession()

	authCode := os.Getenv("TIKTOKADS_AUTH_CODE")
	if len(authCode) == 0 {
		t.Fatal("No auth code received")
	}
	accessToken, err := GetAccessTokenFromAuthCode(authCode)

	if nil != err {
		t.Error(err)
	}

	t.Logf("AccessToken: %s", accessToken)
}
