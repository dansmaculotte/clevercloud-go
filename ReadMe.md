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

## Resources

### [Self](./clevercloud/self.go)

- [API Reference](https://www.clever-cloud.com/doc/api/#!/self)

```go
self, err := clevercloud.GetSelf(client)
selfAddons, err := self.GetAddons(client)
selfAddons, err := self.GetAddons(client, selfAddons[0].ID)
```

### [Organizations](./clevercloud/organization.go)

- [API Reference](https://www.clever-cloud.com/doc/api/#!/organisations)

```go
organizations, err := clevercloud.GetOrganizations(client, self.ID)
organization, err := clevercloud.GetOrganization(client, organizations[0].ID)
organizationAddons, err := organization.GetAddons(client)
organizationAddon, err := organization.GetAddon(client, organizationAddons[0].ID)
```