package gozaim

type MoneySlice struct {
	Money []Money
}

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
