package gozaim

import "encoding/json"

func (c *Client) FetchMe() (Me, error) {
	body, err := c.get("home/user/verify", nil)
	if err != nil {
		return Me{}, err
	}

	var raw struct {
		Me        Me
		Requested string
	}
	json.Unmarshal(body, &raw)

	return raw.Me, nil
}

func (c *Client) FetchMoney(params map[string]string) ([]Money, error) {
	body, err := c.get("home/money", params)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Money []Money
	}
	json.Unmarshal(body, &raw)

	return raw.Money, nil
}

func (c *Client) FetchCategories() ([]Category, error) {
	body, err := c.get("home/category", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Categories []Category
		Requested  int
	}
	json.Unmarshal(body, &raw)

	return raw.Categories, nil
}
