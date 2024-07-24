package models

type Order struct {
	IdOrder   int     `json:"IdOrder"`
	UUIDUser  string  `json:"UUIDUser"`
	IdAddress int     `json:"IdAddress"`
	Total     float64 `json:"Total"`
	CreatedAt string  `json:"CreatedAt"`
}
