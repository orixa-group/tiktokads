package tiktokads

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

// GetOAuthRedirect
func GetOAuthRedirect(redirectUri string) string {
	params := url.Values{}
	params.Set("app_id", appId)
	params.Set("redirect_uri", redirectUri)
	u, _ := url.ParseRequestURI(urlOAuthInstall)
	u.RawQuery = params.Encode()

	return u.String()
}

type accessTokenRequest struct {
	AppId              string `json:"app_id"`
	AppSecret          string `json:"secret"`
	AuthCode           string `json:"auth_code"`
	ReturnAdvertiserId bool   `json:"return_advertiser_id"`
}

type accessTokenResponse struct {
	Data struct {
		AccessToken string `json:"access_token"`
	} `json:"data"`
}

// GetAccessTokenFromAuthCode https://business-api.tiktok.com/portal/docs?id=1739965703387137
func GetAccessTokenFromAuthCode(authCode string) (string, error) {
	u := urlOAuthFetchTokenFromAuthCode

	buf, _ := json.Marshal(&accessTokenRequest{
		AppId:     appId,
		AppSecret: appSecret,
		AuthCode:  authCode,
	})

	req, _ := http.NewRequest("POST", u, bytes.NewBuffer(buf))
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	} else if res.StatusCode >= http.StatusBadRequest {
		return "", errors.New(res.Status)
	}

	buf, _ = io.ReadAll(res.Body)
	token := &accessTokenResponse{}
	if e := json.Unmarshal(buf, token); e != nil {
		return "", e
	} else if len(token.Data.AccessToken) == 0 {
		return "", errors.New("access_token is empty")
	}

	return token.Data.AccessToken, nil
}
