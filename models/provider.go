package models

type Provider struct {
	IdProvider          int    `json:"IdProvider"`
	NameProvider        string `json:"NameProvider"`
	PhoneNumberProvider string `json:"PhoneNumberProvider"`
	EmailProvider       string `json:"EmailProvider"`
}
