package gozaim

import "net/url"

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
