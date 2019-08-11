package gozaim

import (
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"net/http"
	"strings"
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

func (c *Client) get(path string, params map[string]string) ([]byte, error) {
	response, err := c.executeHttpRequest("GET", path, params)
	if err != nil {
		return nil, err
	}
	body, err := parseHttpResponseBody(response)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) executeHttpRequest(method, path string, params map[string]string) (*http.Response, error) {
	request, err := http.NewRequest(method, END_POINT+path+mapToQueryString(params), nil)
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

func mapToQueryString(m map[string]string) string {
	if len(m) == 0 {
		return ""
	}
	queries := []string{}
	for k, v := range m {
		query := k + "=" + v
		queries = append(queries, query)
	}
	return "?" + strings.Join(queries, "&")
}
