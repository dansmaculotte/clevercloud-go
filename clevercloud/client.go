package clevercloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gomodule/oauth1/oauth"
)

// Client Used to request API
type Client struct {
	Crendentials *oauth.Credentials
	OAuthClient  *oauth.Client
	HTTPClient   *http.Client
}

// NewClient Returns a fresh new client instance
func NewClient(config *Config, client *http.Client) *Client {
	return &Client{
		Crendentials: &oauth.Credentials{
			Token:  config.Token,
			Secret: config.Secret,
		},
		OAuthClient: &oauth.Client{
			Credentials: oauth.Credentials{
				Token:  OAuthConsumerKey,
				Secret: OAuthConsumerSecret,
			},
			TemporaryCredentialRequestURI: APIURL + "/oauth/request_token",
			ResourceOwnerAuthorizationURI: APIURL + "/oauth/authorize",
			TokenRequestURI:               APIURL + "/oauth/access_token",
			SignatureMethod:               oauth.HMACSHA1,
		},
		HTTPClient: client,
	}
}

// Get Request with GET method
func (c *Client) Get(uri string, form url.Values, data interface{}) error {
	res, err := c.OAuthClient.Get(c.HTTPClient, c.Crendentials, APIURL+uri, form)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return parseResponse(res, data)
}

// Post Request with POST method
func (c *Client) Post(uri string, form url.Values, data interface{}) error {
	res, err := c.OAuthClient.Post(c.HTTPClient, c.Crendentials, APIURL+uri, form)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return parseResponse(res, data)
}

// Put Request with PUT method
func (c *Client) Put(uri string, form url.Values, data interface{}) error {
	res, err := c.OAuthClient.Put(c.HTTPClient, c.Crendentials, APIURL+uri, form)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return parseResponse(res, data)
}

// Delete Request with DELETE method
func (c *Client) Delete(uri string, form url.Values, data interface{}) error {
	res, err := c.OAuthClient.Delete(c.HTTPClient, c.Crendentials, APIURL+uri, form)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return parseResponse(res, data)
}

func parseResponse(response *http.Response, data interface{}) error {
	if response.StatusCode != 200 {
		message, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("get %s returned status %d, %s", response.Request.URL, response.StatusCode, message)
	}

	return json.NewDecoder(response.Body).Decode(data)
}
