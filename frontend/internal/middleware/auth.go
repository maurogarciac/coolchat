package middleware

import (
	"context"
	"frontend/internal/domain"
	"frontend/internal/services"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenCookieName  = "access_token"
	RefreshTokenCookieName = "refresh_token"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Middleware to validate JWT Access Tokens
func AuthRequired(next http.Handler, b services.BackendService, secretKey string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		// Extract tokens from cookies
		accessToken := GetAccessToken(r)
		refreshToken := GetRefreshToken(r)

		var claims *Claims
		var err error

		if accessToken == "" && refreshToken == "" {
			log.Default().Print("Missing both tokens")
			http.Redirect(w, r, "/login/", http.StatusMovedPermanently)
			return
		}

		if accessToken != "" {
			// Parse and validate the token
			claims, err = VerifyAccessToken(accessToken, secretKey)
			if err != nil {
				log.Default().Print("Invalid or expired token")
				http.Redirect(w, r, "/login/", http.StatusMovedPermanently)
				return
			}
		}

		if accessToken == "" && refreshToken != "" {
			// If access token expired but refresh token didn't, get new access token
			log.Default().Print("Missing refresh token, getting new one")
			newAccessToken, err := b.PostRefresh(ctx, domain.RefreshToken{Token: refreshToken})
			if err != nil {
				log.Default().Print("Error ocurred getting access token from backend")
				http.Redirect(w, r, "/login/", http.StatusMovedPermanently)
			}
			claims, err = VerifyAccessToken(newAccessToken.AccessToken, secretKey)
			if err != nil {
				log.Default().Print("Invalid or expired token")
				// w.Header().Set("HX-Redirect", "/login?partial=true")
				http.Redirect(w, r, "/login/", http.StatusMovedPermanently)
				return
			}

		}

		// Set the username in the request context
		ctx = context.WithValue(ctx, "User", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRefreshToken(r *http.Request) string {

	refreshToken := ""
	for _, cookie := range r.Cookies() {
		if cookie.Name == RefreshTokenCookieName {
			refreshToken = cookie.Value
		}
	}
	return refreshToken
}

func GetAccessToken(r *http.Request) string {

	accessToken := ""
	for _, cookie := range r.Cookies() {
		if cookie.Name == AccessTokenCookieName {
			accessToken = cookie.Value
		}
	}
	return accessToken
}

// Verifies the Access Token
func VerifyAccessToken(tokenStr string, jwtSecretKey string) (*Claims, error) {

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

func SetTokenCookie(name string, token string, expiration time.Time, w http.ResponseWriter) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.SameSite = 0

	http.SetCookie(w, cookie)
}
