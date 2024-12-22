package calculator

import "errors"

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrDivisionByZero    = errors.New("division by zero")
	ErrUnknownToken      = errors.New("unknown token")
	ErrUnknownOperator   = errors.New("unknown operator")
)
