package model

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginatedResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   int64       `json:"total"`
	Page    int         `json:"page"`
	Limit   int         `json:"limit"`
}

// Respon Berhasil
func OK(data interface{}) Response {
	return Response{Success: true, Message: "Success", Data: data}
}

func Created(data interface{}) Response {
	return Response{Success: true, Message: "Created successfully", Data: data}
}

func Updated() Response {
	return Response{Success: true, Message: "Updated successfully"}
}

func Deleted() Response {
	return Response{Success: true, Message: "Deleted successfully"}
}

func OKMessage(message string, data interface{}) Response {
	return Response{Success: true, Message: message, Data: data}
}

// Respon Error
func BadRequest(message string) Response {
	return Response{Success: false, Message: message}
}

func NotFound(resource string) Response {
	return Response{Success: false, Message: resource + " not found"}
}

func InternalError() Response {
	return Response{Success: false, Message: "Something went wrong, please try again later"}
}

func Unauthorized() Response {
	return Response{Success: false, Message: "Unauthorized access"}
}

func Forbidden() Response {
	return Response{Success: false, Message: "You don't have permission to access this resource"}
}

func ValidationError(message string) Response {
	return Response{Success: false, Message: "Validation error: " + message}
}

// Paginated
func OKPaginated(data interface{}, total int64, page, limit int) PaginatedResponse {
	return PaginatedResponse{
		Success: true,
		Message: "Success",
		Data:    data,
		Total:   total,
		Page:    page,
		Limit:   limit,
	}
}