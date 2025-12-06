package domain

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ResponseToken struct {
	Username string `json:"username"`
	Token string `json:"token"`
}