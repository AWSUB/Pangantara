package model

type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type LoginResponse struct {
	Success      bool        `json:"success"`
	Message      string      `json:"message"`
	AccessToken  string      `json:"access_token,omitempty"`
	RefreshToken string      `json:"refresh_token,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func LoginOK(message, accessToken, refreshToken string, data interface{}) LoginResponse {
	return LoginResponse{
		Success:      true,
		Message:      message,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Data:         data,
	}
}

func LoginFail(message string) LoginResponse {
	return LoginResponse{Success: false, Message: message}
}