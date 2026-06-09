package dtos

// LoginDTO carries the payload for a login request.
type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// RefreshTokenDTO carries the payload for a token refresh request.
type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// TokenPairDTO is the response payload containing both tokens.
type TokenPairDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
