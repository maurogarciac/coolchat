package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// Secrets should be in env variables
	AccessTokenCookieName  = "access_token"
	RefreshTokenCookieName = "refresh_token"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Validate Refresh token and return new Access Token
func RefreshAccessToken(
	tokenStr string,
	w http.ResponseWriter,
	r *http.Request,
	jwtSecretKey string,
	jwtRefreshSecretKey string) (string, error) {

	if tokenStr == "" {
		http.Error(w, "Missing refresh_token", http.StatusBadRequest)
		return "", fmt.Errorf("missing refresh_token")
	}

	claims, err := verifyRefreshToken(tokenStr, jwtRefreshSecretKey)
	if err != nil {
		http.Error(w, "Invalid refresh_token", http.StatusBadRequest)
		return "", fmt.Errorf("invalid refresh_token")
	}

	newToken, err := GenerateAccessToken(claims.Username, jwtSecretKey)
	if err != nil {
		http.Error(w, "Could not generate access_token", http.StatusInternalServerError)
		return "", err
	}
	return newToken, nil

}

func verifyRefreshToken(tokenStr string, jwtRefreshSecretKey string) (*Claims, error) {

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtRefreshSecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

// Returns newly generated Access and Refresh Tokens
func GenerateTokens(
	username string,
	jwtSecretKey string,
	jwtRefreshSecretKey string) (string, string, error) {

	accessToken, err := GenerateAccessToken(username, jwtSecretKey)
	if err != nil {
		return "", "", err
	}

	// Refresh Token (24 hours expiration)
	refreshExpiration := time.Now().Add(24 * time.Hour)
	refreshClaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now().Add(10 * time.Millisecond)),
		},
	}
	refreshToken, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(jwtRefreshSecretKey))
	if err != nil {
		return "", "", fmt.Errorf("refresh_token generation error: %w", err)
	}

	return accessToken, refreshToken, nil
}

func GenerateAccessToken(username string, jwtSecretKey string) (string, error) {

	// Access Token (15 minutes expiration)
	accessExpiration := time.Now().Add(15 * time.Minute)
	accessClaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now().Add(10 * time.Millisecond)),
		},
	}

	accessToken, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256, accessClaims).SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", fmt.Errorf("access_token generation error: %w", err)
	}

	return accessToken, nil
}
