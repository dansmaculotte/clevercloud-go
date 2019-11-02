package clevercloud

import (
	"fmt"
	"net/url"

	"github.com/gorilla/schema"
)

var selfBaseURL = "/self"

// Self Resource
type Self struct {
	ID             string   `json:"id"`
	Email          string   `json:"email"`
	Name           string   `json:"name"`
	Phone          string   `json:"phone"`
	Address        string   `json:"address"`
	City           string   `json:"city"`
	ZipCode        string   `json:"zipCode"`
	Country        string   `json:"country"`
	Avatar         string   `json:"avatar"`
	CreationDate   int64    `json:"creationDate"`
	Language       string   `json:"lang"`
	EmailValidated bool     `json:"emailValidated"`
	OAuthApps      []string `json:"oauthApps"`
	Admin          bool     `json:"admin"`
	CanPay         bool     `json:"canPay"`
	PreferredMFA   string   `json:"preferredMFA"`
	HasPassword    bool     `json:"hasPassword"`
}

func GetSelf(client *Client) (*Self, error) {
	self := &Self{}

	if err := client.Get(selfBaseURL, nil, self); err != nil {
		return nil, err
	}

	return self, nil
}

func (s *Self) GetAddons(client *Client) ([]*Addon, error) {
	addons := []*Addon{}

	uri := fmt.Sprintf("%s/%s/addons", selfBaseURL, s.ID)

	if err := client.Get(uri, nil, &addons); err != nil {
		return nil, err
	}

	return addons, nil
}

func (s *Self) GetAddon(client *Client, addonId string) (*Addon, error) {
	addon := &Addon{}

	uri := fmt.Sprintf("%s/%s/addons/%s", selfBaseURL, s.ID, addonId)

	if err := client.Get(uri, nil, &addon); err != nil {
		return nil, err
	}

	return addon, nil
}

type SelfUpdateRequest struct {
	Address  string `schema:"address"`
	City     string `schema:"city"`
	Country  string `schema:"country"`
	Email    string `schema:"email,required"`
	Language string `schema:"language"`
	Name     string `schema:"name"`
	Password string `schema:"password,required"`
	Phone    string `schema:"phone"`
	Terms    bool   `schema:"terms,required"`
	ZipCode  string `schema:"zipcode"`
}

func (s *Self) Update(client *Client, data *SelfUpdateRequest) error {
	encoder := schema.NewEncoder()
	form := url.Values{}

	if err := encoder.Encode(data, form); err != nil {
		return err
	}

	return client.Put(selfBaseURL, form, nil)
}
