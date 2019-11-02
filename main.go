package main

import (
	"fmt"

	"github.com/dansmaculotte/clevercloud-go/clevercloud"
)

func main() {
	config := clevercloud.GetConfig()
	fmt.Println(config)
}
