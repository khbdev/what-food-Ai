package models

type LoginRequest struct {
	Phone string `json:"phone" validate:"required"`
}