package gozaim

// Me ユーザ情報
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
	ProfileImageURL string
	CoverImageURL   string
	ProfileModified string
}
