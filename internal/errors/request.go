package apperror

func ForbiddenError() (int, ErrorResponse) {
	return 401, ErrorResponse{
		Code:    401,
		Message: "forbidden request",
	}
}

func TokenRequiredError() (int, ErrorResponse) {
	return 403, ErrorResponse{
		Code:    403,
		Message: "token required",
	}
}

func NotFound() (int, ErrorResponse) {
	return 404, ErrorResponse{
		Code:    404,
		Message: "not found!",
	}
}

func BadRequest() (int, ErrorResponse) {
	return 400, ErrorResponse{
		Code:    400,
		Message: "Bad request.",
	}
}
