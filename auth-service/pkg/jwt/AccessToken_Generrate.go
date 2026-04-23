package jwt

import (
	"auth-service/internal/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(m models.TokenModel) (string, error) {
	secret := os.Getenv("JWT_ACCESS_SECRET")
	

	expStr := os.Getenv("JWT_ACCESS_EXP_MIN")
	if expStr == "" {
		expStr = "30"
	}

	expMinutes, err := time.ParseDuration(expStr + "m")
	if err != nil {
		expMinutes = 30 * time.Minute
	}

	claims := Claims{
		UserID:   m.UserID,
		UserName: m.UserName,
		Role:     m.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expMinutes)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}