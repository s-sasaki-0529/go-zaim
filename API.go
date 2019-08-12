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

func (c *Client) FetchDefaultCategories() ([]Category, error) {
	body, err := c.get("category", nil)
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

func (c *Client) FetchGenres() ([]Genre, error) {
	body, err := c.get("home/genre", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Genres    []Genre
		Requested int
	}
	json.Unmarshal(body, &raw)

	return raw.Genres, nil
}

func (c *Client) FetchDefaultGenres() ([]Genre, error) {
	body, err := c.get("genre", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Genres    []Genre
		Requested int
	}
	json.Unmarshal(body, &raw)

	return raw.Genres, nil
}

func (c *Client) FetchAccounts() ([]Account, error) {
	body, err := c.get("home/account", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Accounts  []Account
		Requested int
	}
	json.Unmarshal(body, &raw)

	return raw.Accounts, nil
}

func (c *Client) FetchDefaultAccounts() ([]Account, error) {
	body, err := c.get("account", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Accounts  []Account
		Requested int
	}
	json.Unmarshal(body, &raw)

	return raw.Accounts, nil
}

func (c *Client) FetchCurrencies() ([]Currency, error) {
	body, err := c.get("currency", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Currencies  []Currency
		Requested int
	}
	json.Unmarshal(body, &raw)

	return raw.Currencies, nil
}
