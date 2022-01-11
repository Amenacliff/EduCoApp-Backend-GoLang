package dto

type CreateUserResponse struct {
	Success bool        `json:"success"`
	Reason  string      `json:"reason"`
	UserId  interface{} `json:"userId"`
}
