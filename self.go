package clevercloud

import "time"

type Self struct {
	ID             string    `json:"id"`
	Email          string    `json:"email"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	City           string    `json:"city"`
	ZipCode        string    `json:"zipCode"`
	Country        string    `json:"country"`
	Avatar         string    `json:"avatar"`
	CreationDate   time.Time `json:"creationDate"`
	Language       string    `json:"lang"`
	EmailValidated bool      `json:"emailValidated"`
	OAuthApps      []string  `json:"oauthApps"`
	Admin          bool      `json:"admin"`
	CanPay         bool      `json:"canPay"`
	PreferredMFA   string    `json:"preferredMFA"`
	HasPassword    bool      `json:"hasPassword"`
}
