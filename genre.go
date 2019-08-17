package gozaim

// Genre ジャンル
type Genre struct {
	ID            int
	Name          string
	Sort          int
	Active        int
	CategoryID    int
	ParentGenreID int
	Modified      string
}
