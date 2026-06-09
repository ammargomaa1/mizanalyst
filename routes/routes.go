package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mizanalyst/mizanalyst/routes/auth"
)

// RegisterRoutes sets up all application routes on the given Gin engine.
func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		auth.RegisterAuthRoutes(api)
	}

	// Example: protected routes (uncomment and add handlers as needed)
	// protected := api.Group("/")
	// protected.Use(middleware.AuthMiddleware())
	// {
	//     // protected.GET("/profile", profileController.GetProfile)
	// }
}
