package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/mizanalyst/mizanalyst/config"
	"github.com/mizanalyst/mizanalyst/mizanlyst_logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

// GetDB returns the singleton *gorm.DB instance.
// On first call it initialises the connection pool using values from config.
func GetDB() *gorm.DB {
	once.Do(func() {
		cfg := config.GetConfig()

		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			mizanlyst_logger.Log("Failed to connect to database: %v", err)
			panic("failed to connect to database")
		}

		sqlDB, err := db.DB()
		if err != nil {
			mizanlyst_logger.Log("Failed to get underlying sql.DB: %v", err)
			panic("failed to get underlying sql.DB")
		}

		sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConns)
		sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConns)
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.DBConnMaxLifetimeMin) * time.Minute)

		dbInstance = db
		mizanlyst_logger.Log("Database connection established successfully")
	})

	return dbInstance
}
