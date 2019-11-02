package clevercloud

import (
	"fmt"
	"net/url"
)

var organizationsBaseURL = "/organisations"

type Organization struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Address          string `json:"address"`
	City             string `json:"city"`
	ZipCode          string `json:"zipcode"`
	Country          string `json:"country"`
	Company          string `json:"company"`
	VAT              string `json:"VAT"`
	Avatar           string `json:"avatar"`
	VATState         string `json:"vatState"`
	CustomerFullName string `json:"customerFullName"`
	CanPay           bool   `json:"canPay"`
	CleverEnterprise bool   `json:"cleverEnterprise"`
	EmergencyNumber  string `json:"emergencyNumber"`
}

func GetOrganizations(client *Client, ownerId string) ([]*Organization, error) {
	organizations := []*Organization{}
	form := url.Values{}
	form.Add("user", ownerId)

	if err := client.Get(organizationsBaseURL, form, &organizations); err != nil {
		return nil, err
	}

	return organizations, nil
}

func GetOrganization(client *Client, ownerId string) (*Organization, error) {
	organization := &Organization{}

	if err := client.Get(organizationsBaseURL+"/"+ownerId, nil, &organization); err != nil {
		return nil, err
	}

	return organization, nil
}

func (o *Organization) GetAddons(client *Client) ([]*Addon, error) {
	addons := []*Addon{}

	uri := fmt.Sprintf("%s/%s/addons", organizationsBaseURL, o.ID)

	if err := client.Get(uri, nil, &addons); err != nil {
		return nil, err
	}

	return addons, nil
}

func (o *Organization) GetAddon(client *Client, addonId string) (*Addon, error) {
	addon := &Addon{}

	uri := fmt.Sprintf("%s/%s/addons/%s", organizationsBaseURL, o.ID, addonId)

	if err := client.Get(uri, nil, &addon); err != nil {
		return nil, err
	}

	return addon, nil
}
