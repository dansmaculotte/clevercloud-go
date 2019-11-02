package main

import (
	"fmt"

	"github.com/dansmaculotte/clevercloud-go/clevercloud"
)

func main() {
	config := clevercloud.GetConfigFromUser()
	client := clevercloud.NewClient(config, nil)

	self, err := clevercloud.GetSelf(client)
	if err != nil {
		panic(err)
	}

	fmt.Println(self)
}
