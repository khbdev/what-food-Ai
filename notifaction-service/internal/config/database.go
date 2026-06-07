package config

import (
	"fmt"
	"log"
	"os"

	"notifaction-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresConnection() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSLMODE")

	if dbPort == "" {
		dbPort = "5432"
	}

	if dbSSL == "" {
		dbSSL = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Tashkent",
		dbHost, dbUser, dbPass, dbName, dbPort, dbSSL,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("failed to connect PostgreSQL: %v", err)
	}

	// AUTO MIGRATE
	err = db.AutoMigrate(
		&models.Notification{},
	)

	if err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	return db
}
