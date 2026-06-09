package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	bv "github.com/mizanalyst/mizanalyst/business_validators/auth"
	"github.com/mizanalyst/mizanalyst/mizanlyst_logger"
	"github.com/mizanalyst/mizanalyst/requests"
	authReq "github.com/mizanalyst/mizanalyst/requests/auth"
	"github.com/mizanalyst/mizanalyst/responses"
	authSvc "github.com/mizanalyst/mizanalyst/services/auth"
)

// AuthController handles HTTP requests for authentication.
type AuthController struct {
	businessValidator *bv.AuthBusinessValidator
	service           *authSvc.AuthService
}

// NewAuthController creates a new AuthController instance.
func NewAuthController() *AuthController {
	return &AuthController{
		businessValidator: bv.NewAuthBusinessValidator(),
		service:           authSvc.NewAuthService(),
	}
}

// Login handles POST /api/v1/auth/login
func (ctrl *AuthController) Login(c *gin.Context) {
	// 1. Request chain: Bind → Validate
	req := &authReq.LoginRequest{}
	if !req.Run(c) {
		return
	}

	// 2. Business validation
	if err := ctrl.businessValidator.ValidateLogin(req.Body); err != nil {
		responses.BadRequest(c, err.Error())
		return
	}

	// 3. Service call
	tokenPair, err := ctrl.service.Login(req.Body)
	if err != nil {
		if errors.Is(err, authSvc.ErrInvalidCredentials) {
			responses.Unauthorized(c, err.Error())
			return
		}
		mizanlyst_logger.Log("Login error: %v", err)
		responses.BadRequest(c, "An error occurred during login")
		return
	}

	// 4. Success response
	responses.OKWithBody(c, "Login successful", tokenPair)
}

// RefreshToken handles POST /api/v1/auth/refresh
func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	// 1. Request chain: Bind → Validate
	req := &authReq.RefreshTokenRequest{}
	if !req.Run(c) {
		return
	}

	// 2. Business validation
	if err := ctrl.businessValidator.ValidateRefreshToken(req.Body); err != nil {
		responses.BadRequest(c, err.Error())
		return
	}

	// 3. Service call
	tokenPair, err := ctrl.service.RefreshToken(req.Body)
	if err != nil {
		if errors.Is(err, authSvc.ErrInvalidRefreshToken) {
			responses.Unauthorized(c, err.Error())
			return
		}
		mizanlyst_logger.Log("Refresh token error: %v", err)
		responses.BadRequest(c, "An error occurred during token refresh")
		return
	}

	// 4. Success response
	responses.OKWithBody(c, "Token refreshed successfully", tokenPair)
}

// Me handles GET /api/v1/auth/me
func (ctrl *AuthController) Me(c *gin.Context) {
	// 1. Request chain: Retrieve User ID
	var userID uint
	retriever := &requests.RequestUserIdRetriever{UserID: &userID}
	if !requests.RunChain(c, retriever) {
		return
	}

	// 2. Business validation
	if err := ctrl.businessValidator.ValidateMe(userID); err != nil {
		responses.BadRequest(c, err.Error())
		return
	}

	// 3. Service call
	userDTO, err := ctrl.service.Me(userID)
	if err != nil {
		if errors.Is(err, authSvc.ErrUserNotFound) {
			responses.Unauthorized(c, err.Error())
			return
		}
		mizanlyst_logger.Log("Me error: %v", err)
		responses.BadRequest(c, "An error occurred while fetching user details")
		return
	}

	// 4. Success response
	responses.OKWithBody(c, "User details retrieved successfully", userDTO)
}
