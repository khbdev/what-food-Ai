package models



type RegisterRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Age      int    `json:"age" validate:"required,gte=0,lte=120"`
	Address  string `json:"address" validate:"required"`
}