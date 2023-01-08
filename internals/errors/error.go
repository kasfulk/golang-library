package apperror

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func FirebaseLibraryError() (int, ErrorResponse) {
	return 500, ErrorResponse{
		Code:    500,
		Message: "firebase setup error",
	}
}

func ServerError() (int, ErrorResponse) {
	return 500, ErrorResponse{
		Code:    500,
		Message: "Error server",
	}
}

func RedisError() (int, ErrorResponse) {
	return 500, ErrorResponse{
		Code:    500,
		Message: "redis error",
	}
}
