package services

import "errors"

var (
	ErrDailyPostLimitExceeded = errors.New("daily post limit exceeded")
	ErrUserNotFound           = errors.New("user not found")
	ErrInvalidCredentials     = errors.New("invalid credentials")
)
