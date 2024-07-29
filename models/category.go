package models

type Category struct {
	IdCategory          int    `json:"IdCategory"`
	NameCategory        string `json:"NameCategory"`
	DescriptionCategory string `json:"DescriptionCategory"`
	PathCategory        string `json:"PathCategory"`
}
