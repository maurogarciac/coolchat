package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	AccessTokenCookieName  = "access-token"
	JwtSecretKey           = "super-secure-access-token"
	RefreshTokenCookieName = "refresh-token"
	JwtRefreshSecretKey    = "super-secure-secret-key"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Middleware to validate JWT Access Tokens
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Parse and validate the token
		claims, err := verifyAccessToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// You can set the username or other claims in the request context if needed
		ctx := r.Context()
		ctx = context.WithValue(ctx, "username", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Verifies the Access Token
func verifyAccessToken(tokenStr string) (*Claims, error) {

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

// Refresh the Access Token using the Refresh Token
func RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.URL.Query().Get("refresh_token")
	if refreshToken == "" {
		http.Error(w, "Refresh token required", http.StatusUnauthorized)
		return
	}

	claims, err := verifyRefreshToken(refreshToken)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	err = GenerateTokens(claims.Username, w)
	if err != nil {
		http.Error(w, "Could not generate access token", http.StatusInternalServerError)
		return
	}

}

// Verify Refresh Token
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

// Generate Access and Refresh Tokens
func GenerateTokens(username string, w http.ResponseWriter) error {
	// Access Token (15 minutes expiration)
	accessExp := time.Now().Add(15 * time.Minute)
	accessClaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExp),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(JwtSecretKey)
	if err != nil {
		return err
	}

	setTokenCookie(AccessTokenCookieName, accessToken, accessExp, w)

	// Refresh Token (24 hours expiration)
	refreshExp := time.Now().Add(24 * time.Hour)
	refreshClaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExp),
		},
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(JwtRefreshSecretKey)
	if err != nil {
		return err
	}
	setTokenCookie(RefreshTokenCookieName, refreshToken, refreshExp, w)

	return nil
}

func setTokenCookie(name string, token string, expiration time.Time, w http.ResponseWriter) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true

	http.SetCookie(w, cookie)
}
