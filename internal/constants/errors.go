package constants

import "errors"

var (
	ErrIdRequired       = errors.New("error id required")
	ErrEmailExists      = errors.New("error email already exists")
	ErrInvalidAuth      = errors.New("username or password is wrong")
	ErrEmailNotVerified = errors.New("email not verified")
	ErrSessionNotFound  = errors.New("session not found")
	ErrInvalidToken     = errors.New("invalid token")
)
