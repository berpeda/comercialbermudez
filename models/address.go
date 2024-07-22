package models

type Address struct {
	IdAddress         int    `json:"IdAddress"`
	UUIDUser          string `json:"UUIDUser"`
	NameAddress       string `json:"NameAddress"`
	CityAddress       string `json:"CityAddress"`
	StateAddress      string `json:"StateAddress"`
	PhoneAddress      string `json:"PhoneAddress"`
	PostalCodeAddress string `json:"PostalCodeAddress"`
}
