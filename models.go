package gozaim

type Money struct {
	ID            int    `json:"id"`
	Mode          string `json:"mode"`
	UserId        int    `json:"user_id"`
	Date          string `json:"date"`
	CategoryId    int    `json:"category_id"`
	GenreId       int    `json:"genre_id"`
	ToAccountId   int    `json:"to_account_id"`
	FromAccountId int    `json:"from_account_id"`
	Amount        int    `json:"amount"`
	Comment       string `json:"comment"`
	Active        int    `json:"active"`
	Name          string `json:"name"`
	ReceiptId     int    `json:"receipt_id"`
	Place         string `json:"place"`
	Created       string `json:"created"`
	CurrencyCode  int    `json:"currency_code"`
}
