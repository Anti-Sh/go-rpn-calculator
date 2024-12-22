package application

import "errors"

var (
	ErrInvalidRequestMethod = errors.New("invalid request method")
	ErrInternalServerError  = errors.New("internal server error")
)
