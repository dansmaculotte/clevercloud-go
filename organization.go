package clevercloud

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
