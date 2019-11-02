package main

import (
	"fmt"
	"net/http"

	"github.com/dansmaculotte/clevercloud-go/clevercloud"
)

func main() {
	config := clevercloud.GetConfigFromUser()
	client := clevercloud.NewClient(config, &http.Client{})

	self, err := clevercloud.GetSelf(client)
	if err != nil {
		panic(err)
	}

	fmt.Println(self)

	organizations, err := clevercloud.GetOrganizations(client, self.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(organizations)

	organization, err := clevercloud.GetOrganization(client, organizations[0].ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(organization)

	addons, err := organization.GetAddons(client)
	if err != nil {
		panic(err)
	}

	fmt.Println(addons)

	addon, err := organization.GetAddon(client, addons[0].ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(addon)
}
