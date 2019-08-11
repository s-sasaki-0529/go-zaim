package gozaim
import "encoding/json"

func (c *Client) FetchMoney(params map[string]string) (*MoneySlice, error) {
	body, err := c.get("home/money", params)
	if err != nil {
		return nil, err
	}

	var moneySlice MoneySlice
	json.Unmarshal(body, &moneySlice)

	return &moneySlice, nil
}
