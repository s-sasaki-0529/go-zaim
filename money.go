package gozaim

import "net/url"

type MoneySlice []Money

type Money struct {
	ID            int
	Mode          string
	UserId        int
	Date          string
	CategoryId    int
	GenreId       int
	ToAccountId   int
	FromAccountId int
	Amount        int
	Comment       string
	Active        int
	Name          string
	ReceiptId     int
	Place         string
	Created       string
	CurrencyCode  string
}

func (m *Money) Update(client *Client, params url.Values) (bool, error) {
	if m.Mode == "income" {
		return client.UpdateIncome(m.ID, params)
	} else {
		return client.UpdatePayment(m.ID, params)
	}
}

func (m *Money) Delete(client *Client) (bool, error) {
	if m.Mode == "income" {
		return client.DeleteIncome(m.ID)
	} else {
		return client.DeletePayment(m.ID)
	}
}

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
