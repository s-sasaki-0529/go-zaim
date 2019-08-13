package gozaim

type Category struct {
	ID               int
	Name             string
	Mode             string
	Sort             int
	ParentCategoryId int
	Active           int
	Modified         string
}
