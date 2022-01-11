package dto

type LoginRequest struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}
