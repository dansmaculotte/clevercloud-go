package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dansmaculotte/clevercloud-go/clevercloud"
	"github.com/gomodule/oauth1/oauth"
)

var (
	consumerKey    string
	consumerSecret string
	serverAddr     string
)

func readArgs() error {
	flag.StringVar(&consumerKey, "consumer-key", clevercloud.OAuthConsumerKey, "Clever Cloud OAuth consumer key")
	flag.StringVar(&consumerSecret, "consumer-secret", clevercloud.OAuthConsumerSecret, "Clever Cloud OAuth consumer secret")
	flag.StringVar(&serverAddr, "server-addr", "localhost:8000", "Server address for OAuth callback")
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

	client := &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  consumerKey,
			Secret: consumerSecret,
		},
		TemporaryCredentialRequestURI: clevercloud.APIURL + "/oauth/request_token",
		ResourceOwnerAuthorizationURI: clevercloud.APIURL + "/oauth/authorize",
		TokenRequestURI:               clevercloud.APIURL + "/oauth/access_token",
		SignatureMethod:               oauth.HMACSHA1,
	}

	hc := &http.Client{Timeout: 2 * time.Second}

	tmpCredentials, err := client.RequestTemporaryCredentials(hc, "http://"+serverAddr+"/callback", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	authURL := client.AuthorizationURL(tmpCredentials, nil)

	fmt.Printf("Open this URL in your browser:\n%s\n", authURL)

	var verifier string
	var requestToken string
	var userID string

	m := http.NewServeMux()
	s := http.Server{Addr: serverAddr, Handler: m}

	m.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		err = req.ParseForm()

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
		}

		requestToken = req.Form.Get("oauth_token")
		verifier = req.Form.Get("oauth_verifier")
		userID = req.Form.Get("user")

		if requestToken == "" || verifier == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "oauth1: Request missing oauth_token or oauth_verifier")
		}

		w.WriteHeader(200)
		fmt.Fprint(w, "You can close this tab")
		go s.Shutdown(context.Background())
	})

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error listening, %v", err)
	}

	credendials, _, err := client.RequestToken(hc, tmpCredentials, verifier)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("token: %s\nsecret: %s\n", credendials.Token, credendials.Secret)
	fmt.Printf("user: %s\n", userID)
}
