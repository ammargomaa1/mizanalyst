package auth

import (
	"github.com/gin-gonic/gin"
	authCtrl "github.com/mizanalyst/mizanalyst/controllers/auth"
	"github.com/mizanalyst/mizanalyst/middleware"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	authController := authCtrl.NewAuthController()

	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/refresh", authController.RefreshToken)
	}

	// Protected routes
	protected := auth.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/me", authController.Me)
	}
}
