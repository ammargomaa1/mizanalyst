package routes

import (
	"github.com/gin-gonic/gin"
	authCtrl "github.com/mizanalyst/mizanalyst/controllers/auth"
)

// RegisterRoutes sets up all application routes on the given Gin engine.
func RegisterRoutes(router *gin.Engine) {
	authController := authCtrl.NewAuthController()

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/refresh", authController.RefreshToken)
		}
	}

	// Example: protected routes (uncomment and add handlers as needed)
	// protected := api.Group("/")
	// protected.Use(middleware.AuthMiddleware())
	// {
	//     // protected.GET("/profile", profileController.GetProfile)
	// }
}
