package clevercloud

import (
	"net/url"

	"github.com/gorilla/schema"
)

var selfBaseURI = "/self"

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

type SelfAPI struct {
	client *Client
	API
}

func NewSelfAPI(client *Client) *SelfAPI {
	return &SelfAPI{
		client: client,
	}
}

func (s *SelfAPI) GetSelf() (*Self, error) {
	self := &Self{}

	if err := s.client.Get(selfBaseURI, nil, self); err != nil {
		return nil, err
	}

	return self, nil
}

func (s *SelfAPI) GetAddons(ownerID string) ([]*Addon, error) {
	return s.getAddons(s.client, selfBaseURI, ownerID)
}

func (s *SelfAPI) GetAddon(ownerID string, addonID string) (*Addon, error) {
	return s.getAddon(s.client, selfBaseURI, ownerID, addonID)
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

func (s *SelfAPI) UpdateSelf(data *SelfUpdateRequest) error {
	encoder := schema.NewEncoder()
	form := url.Values{}

	if err := encoder.Encode(data, form); err != nil {
		return err
	}

	return s.client.Put(selfBaseURI, form, nil)
}
