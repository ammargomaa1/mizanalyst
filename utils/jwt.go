package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mizanalyst/mizanalyst/config"
)

// Claims represents the JWT claims payload.
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateAccessToken creates a short-lived access token for the given user ID.
func GenerateAccessToken(userID uint) (string, error) {
	cfg := config.GetConfig()

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.AccessTokenTTLMin) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.AccessTokenSecret))
}

// GenerateRefreshToken creates a longer-lived refresh token for the given user ID.
func GenerateRefreshToken(userID uint) (string, error) {
	cfg := config.GetConfig()

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.RefreshTokenTTLDays) * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.RefreshTokenSecret))
}

// ParseAccessToken validates and parses an access token string.
func ParseAccessToken(tokenString string) (*Claims, error) {
	return parseToken(tokenString, config.GetConfig().AccessTokenSecret)
}

// ParseRefreshToken validates and parses a refresh token string.
func ParseRefreshToken(tokenString string) (*Claims, error) {
	return parseToken(tokenString, config.GetConfig().RefreshTokenSecret)
}

// parseToken validates a JWT token against the provided secret.
func parseToken(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
