package gozaim

// Category カテゴリ
type Category struct {
	ID               int
	Name             string
	Mode             string
	Sort             int
	ParentCategoryID int
	Active           int
	Modified         string
}
