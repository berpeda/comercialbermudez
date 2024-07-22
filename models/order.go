package models

type Order struct {
	IdOrder   int    `json:"IdOrder"`
	UUIDUser  string `json:"UUIDUser"`
	IdAddress int    `json:"IdAddress"`
	CreatedAt string `json:"CreatedAt"`
}
