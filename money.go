package gozaim

import "net/url"

// MoneySlice 支払い/収入情報の集合
type MoneySlice []Money

// Money 支払い/収入情報
type Money struct {
	ID            int
	Mode          string
	UserID        int
	Date          string
	CategoryID    int
	GenreID       int
	ToAccountID   int
	FromAccountID int
	Amount        int
	Comment       string
	Active        int
	Name          string
	ReceiptID     int
	Place         string
	Created       string
	CurrencyCode  string
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
