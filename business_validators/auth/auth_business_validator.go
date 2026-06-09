package auth

import "github.com/mizanalyst/mizanalyst/dtos"

// AuthBusinessValidator performs business-level validation for auth operations.
type AuthBusinessValidator struct{}

// NewAuthBusinessValidator creates a new AuthBusinessValidator instance.
func NewAuthBusinessValidator() *AuthBusinessValidator {
	return &AuthBusinessValidator{}
}

// ValidateLogin performs business-rule validations on the login request.
// Currently a pass-through; extend with rules like account lockout, rate limiting, etc.
func (bv *AuthBusinessValidator) ValidateLogin(dto dtos.LoginDTO) error {
	// Add business rules here, e.g.:
	// - Check if account is locked
	// - Check if too many failed attempts
	return nil
}

// ValidateRefreshToken performs business-rule validations on the refresh token request.
// Currently a pass-through; extend with rules like token revocation checks.
func (bv *AuthBusinessValidator) ValidateRefreshToken(dto dtos.RefreshTokenDTO) error {
	// Add business rules here, e.g.:
	// - Check if refresh token has been revoked
	return nil
}

// ValidateMe performs business-rule validations on the me request.
func (bv *AuthBusinessValidator) ValidateMe(userID uint) error {
	return nil
}
