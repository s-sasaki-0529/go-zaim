package gozaim

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/dghubble/oauth1"
)

const END_POINT = "https://api.zaim.net/v2/home/user/verify"

type Client struct {
	HttpClient *http.Client
}

func NewClient(consumerKey, consumerSecret, token, tokenSecret string) *Client {
	oauthConfig := oauth1.NewConfig(consumerKey, consumerSecret)
	oauthToken := oauth1.NewToken(token, tokenSecret)
    httpClient := oauthConfig.Client(oauth1.NoContext, oauthToken)

	resp, err := httpClient.Get(END_POINT)
	if err != nil {
		panic("failed to create client.")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Raw Response Body:\n%v\n", string(body))

	return &Client{HttpClient: httpClient}
}
