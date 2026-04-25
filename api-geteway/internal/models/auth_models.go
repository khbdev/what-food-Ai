package models

// =========================
// AUTH MODELS
// =========================

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Age      int32  `json:"age"`
	Address  string `json:"address"`
}

type LoginRequest struct {
	Phone string `json:"phone"`
}

type VerifyRequest struct {
	OTP int64 `json:"otp"`
}

// =========================
// RESPONSE MODELS
// =========================

type SimpleResponse struct {
	Message string `json:"message"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}