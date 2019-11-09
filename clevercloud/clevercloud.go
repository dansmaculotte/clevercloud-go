package clevercloud

const (
	// APIURL The endpoint for Clever Cloud v2 API
	APIURL = "https://api.clever-cloud.com/v2"

	oAuthConsumerKey    = "x9Dv3WBvHcZZgg3sLmzXTyi2FFhfSu"
	oAuthConsumerSecret = "RnUvL43r4RoUgH9cHqTzeeCh2v0nbv"
)

type Zone string

const (
	Paris    Zone = "par"
	Montreal Zone = "mtl"
)

type API struct {
}
