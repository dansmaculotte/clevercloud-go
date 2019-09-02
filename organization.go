package clevercloud

type Organization struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Avatar           string `json:"avatar"`
	Company          string `json:"company"`
	Address          string `json:"address"`
	zipCode          string `json:"zipcode"`
	City             string `json:"city"`
	Country          string `json:"country"`
	VAT              string `json:"VAT"`
	vatState         string `json:"vatState"`
	customerFullName string `json:"customerFullName"`
}
