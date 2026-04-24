package otp

import (
	"math/rand"
	"time"
)

// GenerateCode 6 xonali OTP code yaratadi
func GenerateCode() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(900000) + 100000
}