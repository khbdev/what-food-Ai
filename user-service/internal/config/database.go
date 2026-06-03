package config

import (
	"fmt"
	"log"	
	"time"
	"user-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB() (*gorm.DB, error) {
	host := "postgres"
	port := "5432"
	user := "postgres"
	password := "secret"
	dbname := "user_service"
	sslmode := "disable"

	if host == "" || port == "" || user == "" || dbname == "" {
		return nil, fmt.Errorf("database environment variables not set properly")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC",
		host, port, user, password, dbname, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn), 
	})
	if err != nil {
		return nil, err
	}

	
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	
	if err := db.AutoMigrate(
		&models.User{},
	); err != nil {
		return nil, err
	}

	log.Println("PostgreSQL connected (GORM) & AutoMigrate done")

	return db, nil
}