package models

import (
	"time"

	"gorm.io/gorm"
)


type Meal struct {
	ID        uint           `gorm:"primaryKey" json:"id"`

	UserID    uint           `gorm:"not null;index" json:"user_id"`

	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Country   string         `gorm:"type:varchar(100)" json:"country"`
    MealTime time.Time `gorm:"not null" json:"meal_time"`

	Kcal      float64        `gorm:"type:decimal(10,2)" json:"kcal"`
	Protein   float64        `gorm:"type:decimal(10,2)" json:"protein"`
	Fat       float64        `gorm:"type:decimal(10,2)" json:"fat"`
	Carbs     float64        `gorm:"type:decimal(10,2)" json:"carbs"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}a