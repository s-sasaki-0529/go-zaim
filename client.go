package gozaim

import (
	"github.com/dghubble/oauth1"
	// "io/ioutil"
	"net/http"
)

const END_POINT = "https://api.zaim.net/v2/"

type Client struct {
	HttpClient *http.Client
}

func NewClient(consumerKey, consumerSecret, token, tokenSecret string) *Client {
	oauthConfig := oauth1.NewConfig(consumerKey, consumerSecret)
	oauthToken := oauth1.NewToken(token, tokenSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, oauthToken)

	return &Client{HttpClient: httpClient}
}

func (c *Client) executeHttpRequest(method, path string) (*http.Response, error) {
	resp, err := c.HttpClient.Get(END_POINT + path)

	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
