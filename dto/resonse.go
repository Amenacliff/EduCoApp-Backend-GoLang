package dto

type CreateUserResponse struct {
	Success bool        `json:"success"`
	Reason  string      `json:"reason"`
	UserId  interface{} `json:"userId"`
}

type LoginUserResponse struct {
	Success      bool        `json:"success"`
	Reason       string      `json:"reason"`
	UserId       interface{} `json:"userId"`
	UserLoggedIn bool        `json:"userLoggedIn"`
}
