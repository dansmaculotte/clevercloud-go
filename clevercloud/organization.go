package clevercloud

import (
	"fmt"
	"net/url"
)

var organizationsBaseURI = "/organisations"

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

type OrganizationAPI struct {
	client *Client
	API
}

func NewOrganizationAPI(client *Client) *OrganizationAPI {
	return &OrganizationAPI{
		client: client,
	}
}

func (o *OrganizationAPI) GetOrganizations(ownerID string) ([]*Organization, error) {
	organizations := []*Organization{}
	form := url.Values{}
	form.Add("user", ownerID)

	if err := o.client.Get(organizationsBaseURI, form, &organizations); err != nil {
		return nil, err
	}

	return organizations, nil
}

func (o *OrganizationAPI) GetOrganization(ownerID string) (*Organization, error) {
	organization := &Organization{}

	uri := fmt.Sprintf("%s/%s", organizationsBaseURI, ownerID)

	if err := o.client.Get(uri, nil, &organization); err != nil {
		return nil, err
	}

	return organization, nil
}

func (o *OrganizationAPI) GetAddons(ownerID string) ([]*Addon, error) {
	return o.getAddons(o.client, organizationsBaseURI, ownerID)
}

func (o *OrganizationAPI) GetAddon(ownerID string, addonID string) (*Addon, error) {
	return o.getAddon(o.client, organizationsBaseURI, ownerID, addonID)
}
