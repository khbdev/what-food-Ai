package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID uint `gorm:"primaryKey"`

	// REQUIRED fields (NOT NULL)
	Name    string `gorm:"size:100;not null"`
	Phone   string `gorm:"size:20;uniqueIndex;not null"`
	Age     int    `gorm:"not null"`
	Address string `gorm:"size:255;not null"`

	// OPTIONAL fields (nullable)
	Email *string `gorm:"size:100;uniqueIndex"`
	Image *string `gorm:"size:255"`

	// Role default user bo‘ladi
	Role Role `gorm:"type:varchar(20);default:user;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time

	// soft delete (GORM style)
	DeletedAt gorm.DeletedAt `gorm:"index"`
}