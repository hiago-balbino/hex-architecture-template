package apperrors

import "errors"

var (
	InternalServerError = errors.New("internal_server_error")
	InvalidInput        = errors.New("invalid_input")
	NotFound            = errors.New("not_found")
)
