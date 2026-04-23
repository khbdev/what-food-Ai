package jwt

import (
	"auth-service/internal/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateRefreshToken(m models.RefreshTokedel) (string, error) {
	secret := os.Getenv("JWT_REFRESH_SECRET")


	expStr := os.Getenv("JWT_REFRESH_EXP_DAYS")
	if expStr == "" {
		expStr = "7"
	}

	expDays, err := time.ParseDuration(expStr + "24h")
	if err != nil {
		expDays = 7 * 24 * time.Hour
	}

	claims := RefreshClaims{
		UserID: m.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expDays)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}