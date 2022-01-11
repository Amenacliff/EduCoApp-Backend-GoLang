package dto

type User struct {
	UserName     string `json:"userName"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}
