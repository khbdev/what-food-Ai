package jwt

import (
	"auth-service/internal/models"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type AccessRefreshClaims struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessRefreshToken(refreshToken string) (string, error) {

	// =========================
	// REFRESH SECRET
	// =========================

	secret := os.Getenv("JWT_REFRESH_SECRET")

	// =========================
	// PARSE REFRESH TOKEN
	// =========================

	token, err := jwt.ParseWithClaims(
		refreshToken,
		&AccessRefreshClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	if err != nil {
		return "", err
	}

	// =========================
	// VALIDATE TOKEN
	// =========================

	claims, ok := token.Claims.(*AccessRefreshClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	// =========================
	// CREATE ACCESS TOKEN
	// =========================

	accessToken, err := GenerateAccessToken(models.TokenModel{
		UserID:   claims.UserID,
		UserName: claims.UserName,
		Role:     claims.Role,
	})

	if err != nil {
		return "", err
	}

	return accessToken, nil
}