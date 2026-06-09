package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mizanalyst/mizanalyst/config"
	"github.com/mizanalyst/mizanalyst/database"
	"github.com/mizanalyst/mizanalyst/migrations"
	"github.com/mizanalyst/mizanalyst/mizanlyst_logger"
	"github.com/mizanalyst/mizanalyst/routes"
	"github.com/mizanalyst/mizanalyst/seeders"
)

func main() {
	// 1. Load configuration
	cfg := config.GetConfig()
	mizanlyst_logger.Log("Configuration loaded successfully")

	// 2. Initialize database connection (singleton)
	db := database.GetDB()
	mizanlyst_logger.Log("Database connection pool initialized")

	// 3. Run migrations
	if err := migrations.RunMigrations(db); err != nil {
		mizanlyst_logger.Log("Migration failed: %v", err)
		panic("failed to run migrations")
	}
	mizanlyst_logger.Log("Migrations completed successfully")

	// 4. Run seeders
	if err := seeders.RunSeeders(db); err != nil {
		mizanlyst_logger.Log("Seeder failed: %v", err)
		panic("failed to run seeders")
	}
	mizanlyst_logger.Log("Seeders completed successfully")

	// 5. Initialize Gin router
	router := gin.Default()

	// 6. Register routes
	routes.RegisterRoutes(router)

	// 7. Start server
	addr := fmt.Sprintf(":%s", cfg.AppPort)
	mizanlyst_logger.Log("Starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		mizanlyst_logger.Log("Failed to start server: %v", err)
		panic("failed to start server")
	}
}
