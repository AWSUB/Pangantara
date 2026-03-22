package model

type CreateUserRequest struct {
	Name     string `json:"name"     binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"     binding:"required,oneof=admin supplier sppg"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"omitempty,email"`
}

type UserResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func UserOK(message string, data interface{}) UserResponse {
	return UserResponse{Success: true, Message: message, Data: data}
}

func UserFail(message string) UserResponse {
	return UserResponse{Success: false, Message: message}
}