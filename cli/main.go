package main

import (
	"errors"
	"flag"
	"os"

	"github.com/dansmaculotte/clevercloud-go"
	"github.com/gomodule/oauth1/oauth"
)

var (
	consumerToken  string
	consumerSecret string
)

var oauthClient = oauth.Client{
	TemporaryCredentialRequestURI: clevercloud.APIURL + "/oauth/request_token",
	ResourceOwnerAuthorizationURI: clevercloud.APIURL + "/oauth/authorize",
	TokenRequestURI:               clevercloud.APIURL + "/oauth/access_token",
	SignatureMethod:               oauth.HMACSHA1,
}

func readArgs() error {
	flag.StringVar(&consumerToken, "consumer-token", "", "Clever Cloud OAuth consumer token")
	flag.StringVar(&consumerSecret, "consumer-secret", "", "Clever Cloud OAuth consumer secret")
	flag.Parse()

	if consumerToken == "" || consumerSecret == "" {
		return errors.New("OAuth Consumer key or secret is required")
	}

	return nil
}

func main() {
	err := readArgs()
	if err != nil {
		os.Exit(1)
	}

	oauthClient.Credentials = oauth.Credentials{
		Token:  consumerToken,
		Secret: consumerSecret,
	}
}
