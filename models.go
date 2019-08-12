package gozaim

type Me struct {
	ID              int
	Login           string
	Name            string
	InputCount      int
	DayCount        int
	RepeatCount     int
	Day             int
	Week            int
	Month           int
	CurrencyCode    string
	ProfileImageUrl string
	CoverImageUrl   string
	ProfileModified string
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

type Category struct {
	ID               int
	Name             string
	Mode             string
	Sort             int
	ParentCategoryId int
	Active           int
	Modified         string
}

type Genre struct {
	ID            int
	Name          string
	Sort          int
	Active        int
	CategoryId    int
	ParentGenreId int
	Modified      int
}
