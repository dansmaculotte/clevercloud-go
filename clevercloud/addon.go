package clevercloud

import (
	"fmt"
)

type Addon struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	RealID       string    `json:"realId"`
	Region       string    `json:"region"`
	CreationDate int64     `json:"creationDate"`
	ConfigKeys   []string  `json:"configKeys"`
	Plan         *Plan     `json:"plan"`
	Provider     *Provider `json:"provider"`
}

type Provider struct {
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	ShortDescription string     `json:"shortDesc"`
	LongDescription  string     `json:"longDesc"`
	LogoURL          string     `json:"logoUrl"`
	OpenInNewTab     bool       `json:"openInNewTab"`
	CanUpgrade       bool       `json:"canUpgrade"`
	Status           string     `json:"status"`
	Regions          []string   `json:"regions"`
	Features         []*Feature `json:"features"`
	AnalyticsID      string     `json:"analyticsId"`
	SupportEmail     string     `json:"supportEmail"`
	Website          string     `json:"website"`
	GooglePlusName   string     `json:"googlePlusName"`
	TwitterName      string     `json:"twitterName"`
}

type Plan struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Slug     string     `json:"slug"`
	Price    float32    `json:"price"`
	Features []*Feature `json:"features"`
}

type Feature struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (a *API) getAddons(client *Client, baseURL, ownerID string) ([]*Addon, error) {
	addons := []*Addon{}

	uri := fmt.Sprintf("%s/%s/addons", baseURL, ownerID)

	if err := client.Get(uri, nil, &addons); err != nil {
		return nil, err
	}

	return addons, nil
}

func (a *API) getAddon(client *Client, baseURL, ownerID, addonID string) (*Addon, error) {
	addon := &Addon{}

	uri := fmt.Sprintf("%s/%s/addons/%s", baseURL, ownerID, addonID)

	if err := client.Get(uri, nil, &addon); err != nil {
		return nil, err
	}

	return addon, nil
}

func (a *API) getAddonProviders(client *Client, baseURL, ownerID string) ([]*Addon, error) {
	addons := []*Addon{}

	uri := fmt.Sprintf("%s/%s/addonproviders", baseURL, ownerID)

	if err := client.Get(uri, nil, &addons); err != nil {
		return nil, err
	}

	return addons, nil
}

func (a *API) getAddonProvider(client *Client, baseURL, ownerID, addonID string) (*Addon, error) {
	addon := &Addon{}

	uri := fmt.Sprintf("%s/%s/addonproviders/%s", baseURL, ownerID, addonID)

	if err := client.Get(uri, nil, &addon); err != nil {
		return nil, err
	}

	return addon, nil
}
