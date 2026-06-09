package auth
import (
	"github.com/gin-gonic/gin"
	authCtrl "github.com/mizanalyst/mizanalyst/controllers/auth"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	authController := authCtrl.NewAuthController()

	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/refresh", authController.RefreshToken)
	}

	// Example: protected routes (uncomment and add handlers as needed)
	// protected := api.Group("/")
	// protected.Use(middleware.AuthMiddleware())
	// {
	//     // protected.GET("/profile", profileController.GetProfile)
	// }
}