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
	JwtSecretKey           = "Bs2S9WLytsE8nPjIMzbd3FgE6VAVODTh"
	RefreshTokenCookieName = "refresh_token"
	JwtRefreshSecretKey    = "hWJH1THzDv03eYjIswvqz6cSqKnpuKrw"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Validate Refresh token and return new Access Token
func RefreshAccessToken(tokenStr string, w http.ResponseWriter, r *http.Request) (string, error) {

	if tokenStr == "" {
		http.Error(w, "Missing refresh token", http.StatusUnauthorized)
		return "", fmt.Errorf("missing refresh token")
	}

	claims, err := verifyRefreshToken(tokenStr)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return "", fmt.Errorf("invalid refresh token")
	}

	newToken, err := GenerateAccessToken(claims.Username)
	if err != nil {
		http.Error(w, "Could not generate access token", http.StatusInternalServerError)
		return "", err
	}
	return newToken, nil

}

func verifyRefreshToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtRefreshSecretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

// Returns newly generated Access and Refresh Tokens
func GenerateTokens(username string) (string, string, error) {

	accessToken, err := GenerateAccessToken(username)
	if err != nil {
		return "", "", err
	}
	// Refresh Token (24 hours expiration)
	refreshExpiration := time.Now().Add(24 * time.Hour)
	refreshClaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpiration),
		},
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(JwtRefreshSecretKey))
	if err != nil {
		return "", "", fmt.Errorf("refresh token generation error: %w", err)
	}

	return accessToken, refreshToken, nil
}

func GenerateAccessToken(username string) (string, error) {
	// Access Token (15 minutes expiration)
	accessExpiration := time.Now().Add(15 * time.Minute)
	accessClaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiration),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(JwtSecretKey))
	if err != nil {
		return "", fmt.Errorf("access token generation error: %w", err)
	}

	return accessToken, nil
}