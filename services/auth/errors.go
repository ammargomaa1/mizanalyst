package auth

import "errors"

// Sentinel errors for the auth service layer.
var (
	ErrInvalidCredentials  = errors.New("invalid email or password")
	ErrInvalidRefreshToken = errors.New("invalid or expired refresh token")
	ErrUserNotFound = errors.New("user not found")
)
