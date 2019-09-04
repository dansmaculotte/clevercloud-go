package clevercloud

import (
	"net/http"

	"github.com/dghubble/oauth1"
)

// APIURL is the endpoint for Clever Cloud v2 API
var APIURL = "https://api.clever-cloud.com/v2"

// NewClient return an HTTP Client with OAuth headers
func NewClient(accessToken string, accessSecret string) *http.Client {
	config := &oauth1.Config{
		ConsumerKey:    "",
		ConsumerSecret: "",
		Endpoint: oauth1.Endpoint{
			RequestTokenURL: APIURL + "/oauth/request_token",
			AuthorizeURL:    APIURL + "/oauth/authorize",
			AccessTokenURL:  APIURL + "/oauth/access_token",
		},
		Realm: APIURL + "/oauth",
	}
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	return httpClient
}
