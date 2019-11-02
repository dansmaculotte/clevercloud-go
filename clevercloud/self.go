package clevercloud

var BaseUrl = "/self"

// Self https://www.clever-cloud.com/doc/api/#!/self
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

	if err := client.Get(BaseUrl, nil, self); err != nil {
		return nil, err
	}

	return self, nil
}
