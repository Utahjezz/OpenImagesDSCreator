package model

type Category struct {
	LabelName   string     `json:"LabelName"`
	SubCategory []Category `json:"Subcategory"`
}
