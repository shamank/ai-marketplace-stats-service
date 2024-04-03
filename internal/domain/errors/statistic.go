package apperrors

import "errors"

var (
	ErrNotFound        = errors.New("statistic not found")
	ErrUserNotFound    = errors.New("user not found")
	ErrServiceNotFound = errors.New("service not found")
)
