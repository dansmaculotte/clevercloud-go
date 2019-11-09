package clevercloud

import (
	"fmt"
)

type Application struct {
	ID             string      `json:"id"`
	OwnerID        string      `json:"ownerId"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	Branch         string      `json:"branch"`
	State          string      `json:"state"`
	Zone           Zone        `json:"zone"`
	StickySessions bool        `json:"stickySessions"`
	SeperateBuild  bool        `json:"separateBuild"`
	Archived       bool        `json:"archived"`
	CancelOnPush   bool        `json:"cancelOnPush"`
	BuildFlavor    *Flavor     `json:"buildFlavor"`
	CommitID       string      `json:"commitId"`
	LastDeploy     int         `json:"lastDeploy"`
	CreationDate   int         `json:"creationDate"`
	DeployURL      string      `json:"deployUrl"`
	Deployment     *Deployment `json:"deployment"`
	Favourite      bool        `json:"favourite"`
	Homogeneous    bool        `json:"homogeneous"`
	VHosts         []*VHost    `json:"vhosts"`
	WebhookURL     string      `json:"webhookUrl"`
	WebhookSecret  string      `json:"webhookSecret"`
}

type Flavor struct {
	Name         string  `json:"name"`
	Available    bool    `json:"available,omitempty"`
	CPUs         int     `json:"cpus"`
	Memory       int     `json:"mem"`
	Disk         int     `json:"disk"`
	Microservice bool    `json:"microservice"`
	Price        float32 `json:"price"`
	Nice         int     `json:"nice"`
}

type Deployment struct {
	URL          string `json:"url"`
	HTTPURL      string `json:"httpUrl,omitempty"`
	Type         string `json:"type"`
	RepoState    string `json:"repoState"`
	Shutdownable bool   `json:"shutdownable"`
}

type Instance struct {
	Name                string            `json:"name,omitempty"`
	Description         string            `json:"description,omitempty"`
	Type                string            `json:"type"`
	Variant             *Variant          `json:"variant"`
	Version             string            `json:"version"`
	Tags                []string          `json:"tags,omitempty"`
	ComingSoon          bool              `json:"comingSoon,omitempty"`
	Enabled             bool              `json:"enabled,omitempty"`
	DefaultEnv          map[string]string `json:"defaultEnv"`
	Deployments         []string          `json:"deployments,omitempty"`
	BuildFlavor         *Flavor           `json:"buildFlavor,omitempty"`
	DefaultFlavor       *Flavor           `json:"defaultFlavor,omitempty"`
	Flavors             []*Flavor         `json:"flavors"`
	InstanceAndVersion  string            `json:"instanceAndVersion"`
	MaxAllowedInstances int               `json:"maxAllowedInstances"`
	MaxFlavor           *Flavor           `json:"maxFlavor"`
	MaxInstances        int               `json:"maxInstances"`
	MinInstances        int               `json:"minInstances"`
}

type Variant struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Logo       string `json:"logo"`
	DeployType string `json:"deployType"`
}

type VHost struct {
	FQDN string `json:"fqdn"`
}

func getApplications(client *Client, baseURI, ownerID string) ([]*Application, error) {
	applications := []*Application{}

	uri := fmt.Sprintf("%s/%s/applications", baseURI, ownerID)

	if err := client.Get(uri, nil, &applications); err != nil {
		return nil, err
	}

	return applications, nil
}

func getApplication(client *Client, baseURI, ownerID, applicationID string) (*Application, error) {
	application := &Application{}

	uri := fmt.Sprintf("%s/%s/applications/%s", baseURI, ownerID, applicationID)

	if err := client.Get(uri, nil, &application); err != nil {
		return nil, err
	}

	return application, nil
}
