# clevercloud-go

> A Go library for interacting with [Clever Cloud's API](https://www.clever-cloud.com/doc/clever-cloud-apis/cc-api/).

## Installation

```bash
go get github.com/dansmaculotte/clevercloud-go/clevercloud
```

## OAuth

Generate an `accessToken` and `accessSecret` with [Clever Cloud CLI OAuth](https://console.clever-cloud.com/cli-oauth).

### via CLI

Install Go 1.13+ and run the following command :

```bash
go run github.com/dansmaculotte/clevercloud-go/cli
```

## Usage

To properly request a resource you need a fresh new client instance with configuration :

```go
import "github.com/dansmaculotte/clevercloud-go/clevercloud"

config := &clevercloud.Config{
    Token: "mytoken",
    Secret: "mysecret"
}
// or get config from user clever-cloud file usually in $HOME/.config/clever-cloud
config := clevercloud.GetConfigFromUser()

client := clevercloud.NewClient(config, httpClient)
```