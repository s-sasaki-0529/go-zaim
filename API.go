package gozaim

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// FetchMe 認証中ユーザ情報をフェッチ
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

// FetchMoney 収入、支出の一覧をフェッチ
func (c *Client) FetchMoney(params url.Values) (MoneySlice, error) {
	body, err := c.get("home/money", params)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Money MoneySlice
	}
	json.Unmarshal(body, &raw)

	return raw.Money, nil
}

// FetchCategories カテゴリ一覧をフェッチ
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

// FetchDefaultCategories デフォルトのカテゴリ一覧をフェッチ
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

// FetchGenres ジャンル一覧をフェッチ
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

// FetchDefaultGenres デフォルトのジャンル一覧をフェッチ
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

// FetchAccounts 口座一覧をフェッチ
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

// FetchDefaultAccounts デフォルトの口座一覧をフェッチ
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

// FetchCurrencies 通貨一覧をフェッチ
func (c *Client) FetchCurrencies() ([]Currency, error) {
	body, err := c.get("currency", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Currencies []Currency
		Requested  int
	}
	json.Unmarshal(body, &raw)

	return raw.Currencies, nil
}

// CreatePayment 支払い情報をPOST
func (c *Client) CreatePayment(params url.Values) (bool, error) {
	_, err := c.post("home/money/payment", params)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateIncome 収入情報をPOST
func (c *Client) CreateIncome(params url.Values) (bool, error) {
	_, err := c.post("home/money/income", params)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdatePayment 支払情報をUPDATE
func (c *Client) UpdatePayment(id int, params url.Values) (bool, error) {
	_, err := c.put("home/money/payment/"+strconv.Itoa(id), params)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateIncome 収入情報をUPDATE
func (c *Client) UpdateIncome(id int, params url.Values) (bool, error) {
	_, err := c.put("home/money/income/"+strconv.Itoa(id), params)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeletePayment 支払情報DELETE
func (c *Client) DeletePayment(id int) (bool, error) {
	_, err := c.put("home/money/payment/"+strconv.Itoa(id), nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteIncome 収入情報をDELETE
func (c *Client) DeleteIncome(id int) (bool, error) {
	_, err := c.put("home/money/income/"+strconv.Itoa(id), nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
