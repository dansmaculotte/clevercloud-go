package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	clevercloud "github.com/dansmaculotte/clevercloud-go"
	"github.com/dghubble/oauth1"
)

var (
	consumerKey    string
	consumerSecret string
)

var config *oauth1.Config

func readArgs() error {
	flag.StringVar(&consumerKey, "consumer-key", "", "Clever Cloud OAuth consumer key")
	flag.StringVar(&consumerSecret, "consumer-secret", "", "Clever Cloud OAuth consumer secret")
	flag.Parse()

	if consumerKey == "" || consumerSecret == "" {
		return errors.New("OAuth Consumer key or secret are required")
	}

	return nil
}

func main() {
	err := readArgs()
	if err != nil {
		os.Exit(1)
	}

	config = &oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    "http://localhost:8080/callback",
		Endpoint: oauth1.Endpoint{
			RequestTokenURL: clevercloud.APIURL + "/oauth/request_token",
			AuthorizeURL:    clevercloud.APIURL + "/oauth/authorize",
			AccessTokenURL:  clevercloud.APIURL + "/oauth/access_token",
		},
		Realm: clevercloud.APIURL + "/oauth",
	}

	requestToken, requestSecret, err := login()
	if err != nil {
		log.Fatalf("Request Token Phase: %s", err.Error())
	}

	accessToken, err := access(requestToken, requestSecret)
	if err != nil {
		log.Fatalf("Access Token Phase: %s", err.Error())
	}

	fmt.Println("Consumer was granted an access token to act on behalf of a user.")
	fmt.Printf("token: %s\nsecret: %s\n", accessToken.Token, accessToken.TokenSecret)
}

func login() (requestToken, requestSecret string, err error) {
	requestToken, requestSecret, err = config.RequestToken()
	if err != nil {
		return "", "", err
	}
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		return "", "", err
	}
	fmt.Printf("Open this URL in your browser:\n%s\n", authorizationURL.String())
	return requestToken, requestSecret, err
}

func access(requestToken, requestSecret string) (*oauth1.Token, error) {
	var verifier string
	var err error

	m := http.NewServeMux()
	s := http.Server{Addr: ":8080", Handler: m}

	m.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		requestToken, verifier, err = oauth1.ParseAuthorizationCallback(req)
		s.Shutdown(context.Background())
	})
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Error listening, %v", err)
	}

	accessToken, accessSecret, err := config.AccessToken(requestToken, requestSecret, verifier)
	if err != nil {
		return nil, err
	}
	return oauth1.NewToken(accessToken, accessSecret), err
}
