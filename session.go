package tiktokads

var appId = ""
var appSecret = ""
var accessToken = ""

// SetApiKeys Set app_id and app_secret
// Used for oauth process
func SetApiKeys(id, secret string) {
	appId = id
	appSecret = secret
}

// SetAccessToken Set access token
func SetAccessToken(token string) {
	accessToken = token
}
