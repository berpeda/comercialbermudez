package models

type User struct {
	UUIDUser    string `json:"UUIDUser"`
	NameUser    string `json:"NameUser"`
	SurnameUser string `json:"SurnameUser"`
	EmailUser   string `json:"EmailUser"`
	RolUser     int    `json:"RolUser"`
	CreatedAt   string `json:"CreatedAt"`
}
