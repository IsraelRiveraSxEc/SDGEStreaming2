package errors

type AppError struct {
	Code string `json:"code"`
	Message string `json:"message"`
	HTTPStatus int `json:"http_status"`
}

