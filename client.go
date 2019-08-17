package gozaim

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dghubble/oauth1"
)

// EndPoint APIエンドポイントのURL
const EndPoint = "https://api.zaim.net/v2/"

// Client OAuthクライアント
type Client struct {
	HTTPClient *http.Client
}

// NewClient OAuthクライアントインスタンスを生成
func NewClient(consumerKey, consumerSecret, token, tokenSecret string) *Client {
	oauthConfig := oauth1.NewConfig(consumerKey, consumerSecret)
	oauthToken := oauth1.NewToken(token, tokenSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, oauthToken)

	return &Client{HTTPClient: httpClient}
}

func (c *Client) get(path string, params url.Values) ([]byte, error) {
	return c.executeHTTPRequestAndParseBody("GET", path, params)
}

func (c *Client) post(path string, params url.Values) ([]byte, error) {
	return c.executeHTTPRequestAndParseBody("POST", path, params)
}

func (c *Client) put(path string, params url.Values) ([]byte, error) {
	return c.executeHTTPRequestAndParseBody("PUT", path, params)
}

func (c *Client) delete(path string, params url.Values) ([]byte, error) {
	return c.executeHTTPRequestAndParseBody("delete", path, params)
}

func (c *Client) executeHTTPRequestAndParseBody(method, path string, params url.Values) ([]byte, error) {
	response, err := c.executeHTTPRequest(method, path, params)
	if err != nil {
		return nil, err
	}
	body, err := parseHTTPResponseBody(response)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	return body, nil
}

func (c *Client) executeHTTPRequest(method, path string, params url.Values) (*http.Response, error) {
	var request *http.Request
	var err error

	if method == "GET" {
		request, err = http.NewRequest(
			method,
			EndPoint+path+"?"+params.Encode(),
			nil,
		)
	} else {
		request, err = http.NewRequest(
			method,
			EndPoint+path,
			bytes.NewBufferString(params.Encode()),
		)
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	if err != nil {
		return nil, err
	}
	return c.HTTPClient.Do(request)
}

func parseHTTPResponseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
