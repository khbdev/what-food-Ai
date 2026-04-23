package models

type AuthMessage struct {
	Task  string `json:"task"`
	Phone string `json:"phone"`
	OTP   string `json:"otp"`
}