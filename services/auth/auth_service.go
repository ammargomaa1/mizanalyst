package auth

import (
	"github.com/mizanalyst/mizanalyst/dtos"
	userRepo "github.com/mizanalyst/mizanalyst/repositories/user"
	"github.com/mizanalyst/mizanalyst/utils"
)

// AuthService handles authentication business logic.
type AuthService struct {
	userRepository *userRepo.UserRepository
}

// NewAuthService creates a new AuthService instance.
func NewAuthService() *AuthService {
	return &AuthService{
		userRepository: userRepo.NewUserRepository(),
	}
}

// Login authenticates a user and returns a token pair.
func (s *AuthService) Login(dto dtos.LoginDTO) (*dtos.TokenPairDTO, error) {
	user, err := s.userRepository.FindByEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidCredentials
	}

	if !utils.CheckPassword(dto.Password, user.Password) {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &dtos.TokenPairDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// RefreshToken validates a refresh token and issues a new token pair.
func (s *AuthService) RefreshToken(dto dtos.RefreshTokenDTO) (*dtos.TokenPairDTO, error) {
	claims, err := utils.ParseRefreshToken(dto.RefreshToken)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	// Verify the user still exists
	user, err := s.userRepository.FindByID(claims.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidRefreshToken
	}

	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &dtos.TokenPairDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Me retrieves the authenticated user's details.
func (s *AuthService) Me(userID uint) (*dtos.UserDTO, error) {
	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return &dtos.UserDTO{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}
