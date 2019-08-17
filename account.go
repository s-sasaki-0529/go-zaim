package gozaim

// Account 口座情報
type Account struct {
	ID              int
	Name            string
	Mode            string
	Sort            int
	ParentAccountID int
	Active          int
	Modified        string
}
