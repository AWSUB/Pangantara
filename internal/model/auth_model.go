package model

type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Token   string      `json:"token,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func LoginOK(message, token string, data interface{}) LoginResponse {
	return LoginResponse{Success: true, Message: message, Token: token, Data: data}
}

func LoginFail(message string) LoginResponse {
	return LoginResponse{Success: false, Message: message}
}