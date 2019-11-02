package clevercloud

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
