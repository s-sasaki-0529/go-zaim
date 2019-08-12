package gozaim

import "encoding/json"

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
