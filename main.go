package main

import (
	"fmt"
	"net/http"

	"github.com/dansmaculotte/clevercloud-go/clevercloud"
)

func main() {
	config := clevercloud.GetConfigFromUser()
	client := clevercloud.NewClient(config, &http.Client{})

	fmt.Println(client)

	selfAPI := clevercloud.NewSelfAPI(client)

	self, err := selfAPI.GetSelf()
	if err != nil {
		panic(err)
	}

	fmt.Print(self)
}
