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

	
	Name    string `gorm:"size:100;not null"`
	Phone   string `gorm:"size:20;uniqueIndex;not null"`
	Age     int    `gorm:"not null"`
	Address string `gorm:"size:255;not null"`

Email string `gorm:"uniqueIndex,default:null"`
	Image string `gorm:"size:255"`

	
	Role Role `gorm:"type:varchar(20);default:user;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"index"`
}