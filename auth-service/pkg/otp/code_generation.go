package otp

import (
	"math/rand"
	"time"
)

func GenerateCode() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(900000) + 100000
}