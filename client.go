package gozaim

import (
	"bytes"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"net/http"
	"net/url"
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

func (c *Client) get(path string, params url.Values) ([]byte, error) {
	return c.executeHttpRequestAndParseBody("GET", path, params)
}

func (c *Client) Post(path string, params url.Values) ([]byte, error) {
	return c.executeHttpRequestAndParseBody("POST", path, params)
}

func (c *Client) executeHttpRequestAndParseBody(method, path string, params url.Values) ([]byte, error) {
	response, err := c.executeHttpRequest(method, path, params)
	if err != nil {
		return nil, err
	}
	body, err := parseHttpResponseBody(response)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) executeHttpRequest(method, path string, params url.Values) (*http.Response, error) {
	var request *http.Request
	var err error

	if method == "GET" {
		request, err = http.NewRequest(
			method,
			END_POINT+path+"?"+params.Encode(),
			nil,
		)
	} else {
		request, err = http.NewRequest(
			method,
			END_POINT+path,
			bytes.NewBufferString(params.Encode()),
		)
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	if err != nil {
		return nil, err
	}
	return c.HttpClient.Do(request)
}

func parseHttpResponseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else {
		return body, nil
	}
}
