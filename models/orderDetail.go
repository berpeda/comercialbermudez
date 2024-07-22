package models

type OrderDetails struct {
	IdOrderDetail       int     `json:"IdOrderDetail"`
	IdOrder             int     `json:"IdOrder"`
	IdProduct           int     `json:"IdProduct"`
	QuantityOrderDetail int     `json:"QuantityOrderDetail"`
	PriceOrderDetail    float64 `json:"PriceOrderDetail"`
}
