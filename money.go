package gozaim

import "net/url"

// MoneySlice 支払い/収入情報の集合
type MoneySlice []Money

// Money 支払い/収入情報
type Money struct {
	ID            int    `json:"id"`
	Mode          string `json:"mode"`
	UserID        int    `json:"user_id"`
	Date          string `json:"date"`
	CategoryID    int    `json:"category_id"`
	GenreID       int    `json:"genre_id"`
	ToAccountID   int    `json:"to_account_id"`
	FromAccountID int    `json:"from_account_id"`
	Amount        int    `json:"amount"`
	Comment       string `json:"comment"`
	Active        int    `json:"active"`
	Name          string `json:"name"`
	ReceiptID     int    `json:"receipt_id"`
	Place         string `json:"place"`
	Created       string `json:"created"`
	CurrencyCode  string `json:"currency_code"`
}

// Update 支払い/収入情報を更新
func (m *Money) Update(client *Client, params url.Values) (bool, error) {
	if m.Mode == "income" {
		return client.UpdateIncome(m.ID, params)
	}
	return client.UpdatePayment(m.ID, params)
}

// Delete 支払い/収入情報を削除
func (m *Money) Delete(client *Client) (bool, error) {
	if m.Mode == "income" {
		return client.DeleteIncome(m.ID)
	}
	return client.DeletePayment(m.ID)
}

// UpdateAll 複数の支払い/収入情報を更新
func (ms MoneySlice) UpdateAll(client *Client, params url.Values) ([]bool, []error) {
	var results []bool
	var errors []error

	for _, m := range ms {
		r, e := m.Update(client, params)
		results = append(results, r)
		errors = append(errors, e)
	}
	return results, errors
}

// DeleteAll 複数の支払い/収入情報を削除
func (ms MoneySlice) DeleteAll(client *Client) ([]bool, []error) {
	var results []bool
	var errors []error

	for _, m := range ms {
		r, e := m.Delete(client)
		results = append(results, r)
		errors = append(errors, e)
	}
	return results, errors
}
