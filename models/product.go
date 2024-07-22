package models

type Product struct {
	IdProduct          int     `json:"IdProduct"`
	IdProvider         int     `json:"IdProvider"`
	IdCategory         int     `json:"IdCategory"`
	CodeProduct        string  `json:"CodeProduct"`
	NameProduct        string  `json:"NameProduct"`
	DescriptionProduct string  `json:"DescriptionProduct"`
	PriceProduct       float64 `json:"PriceProduct"`
	CreatedAt          string  `json:"CreatedAt"`
	UpdatedAt          string  `json:"UpdatedAt"`
	Stock              int     `json:"Stock"`
}
