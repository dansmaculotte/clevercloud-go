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

### [Self](https://www.clever-cloud.com/doc/api/#!/self)

```go
self, err := clevercloud.GetSelf(client)
```

#### [Addons](https://www.clever-cloud.com/doc/api/#!/addons)

```go
addons, err := self.GetAddons(client)
addon, err := self.GetAddons(client, addon.ID)
```

### [Organizations](https://www.clever-cloud.com/doc/api/#!/organisations)

```go
organizations, err := clevercloud.GetOrganizations(client, "owner_id")
organization, err := clevercloud.GetOrganization(client, organizations[0].ID)
```

#### [Addons](https://www.clever-cloud.com/doc/api/#!/addons)

```go
addons, err := organization.GetAddons(client)
addon, err := organization.GetAddon(client, addon.ID)
```