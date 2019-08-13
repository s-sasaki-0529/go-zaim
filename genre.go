package gozaim

type Genre struct {
	ID            int
	Name          string
	Sort          int
	Active        int
	CategoryId    int
	ParentGenreId int
	Modified      string
}
