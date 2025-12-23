package database

import (
	"fmt"
	"time"

	"comment/src/config"
	"comment/src/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect initializes database connection using GORM
func Connect(dbHost, dbName string) *gorm.DB {
	// 🔹 Database connection string (DSN)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Dhaka",
		dbHost,
		config.DBUser,
		config.DBPassword,
		dbName,
		config.DBPort,
	)

	// 🔹 Open database using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info), // show SQL logs
		SkipDefaultTransaction: true,                                // better performance
		PrepareStmt:            true,                                // prepared statements
		TranslateError:         true,                                // db error → gorm error
	})
	if err != nil {
		utils.Log.Fatalf("❌ Failed to connect to database: %+v", err)
	}

	// 🔹 Get underlying sql.DB to configure pool
	sqlDB, err := db.DB()
	if err != nil {
		utils.Log.Fatalf("❌ Failed to get database instance: %+v", err)
	}

	// 🔹 Connection pool configuration
	sqlDB.SetMaxIdleConns(10)                  // idle connections
	sqlDB.SetMaxOpenConns(100)                 // max open connections
	sqlDB.SetConnMaxLifetime(60 * time.Minute) // connection lifetime

	utils.Log.Info("✅ Database connected successfully")

	return db
}
