# clevercloud-go

> A Go library for interacting with [Clever Cloud's API](https://www.clever-cloud.com/doc/clever-cloud-apis/cc-api/).

## Installation

```bash
go get github.com/dansmaculotte/clevercloud-go/clevercloud
```

## OAuth

Generate an `accessToken` and `accessSecret` with [Clever Cloud CLI OAuth](https://console.clever-cloud.com/cli-oauth).

## Usage

To properly request a resource you need a fresh new client instance:

```go
import "github.com/dansmaculotte/clevercloud-go/clevercloud"

config := &clevercloud.Config{
    Token: "mytoken",
    Secret: "mysecret"
}
// or get config from $HOME/.config/clever-cloud
config = clevercloud.GetConfigFromUser()
// or get config from env with CLEVER_TOKEN and CLEVER_SECRET
config = clevercloud.GetConfigFromEnv()

client := clevercloud.NewClient(config, &http.Client{})
```