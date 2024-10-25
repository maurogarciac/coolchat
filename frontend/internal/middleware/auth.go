package middleware

import (
	"context"
	"frontend/internal/domain"
	"frontend/internal/services"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

const (
	AccessTokenCookieName  = "access_token"
	RefreshTokenCookieName = "refresh_token"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AuthStuff struct {
	Lg        *zap.SugaredLogger
	Back      services.BackendService
	SecretKey string
}

// Middleware to validate JWT Access Tokens
func AuthRequired(next http.Handler, a AuthStuff) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		a.Lg.Info("Going through the auth middleware")

		redirect_link := "/login/?partial=true"

		referer := r.Header.Get("Referer")

		if referer == "" || !strings.Contains(referer, "localhost:8000") {
			redirect_link = "/login/"
		}

		ctx := r.Context()

		// Extract tokens from cookies
		accessToken := GetAccessToken(r)
		refreshToken := GetRefreshToken(r)

		var claims *Claims
		var err error

		if accessToken == "" && refreshToken == "" {
			a.Lg.Info("Missing both tokens")
			http.Redirect(w, r, redirect_link, http.StatusMovedPermanently)
			return
		}

		if accessToken != "" {
			// Parse and validate the token
			claims, err = VerifyAccessToken(accessToken, a.SecretKey)
			if err != nil {
				a.Lg.Info("Invalid or expired token")
				http.Redirect(w, r, redirect_link, http.StatusMovedPermanently)
				return
			}
		}

		if accessToken == "" && refreshToken != "" {
			// If access token expired but refresh token didn't, get new access token
			a.Lg.Info("Missing access token, refreshing")

			newAccessToken, err := a.Back.PostRefresh(ctx, domain.RefreshToken{Token: refreshToken})
			if err != nil {
				a.Lg.Info("Error ocurred getting access token from backend: %s", err)
				http.Redirect(w, r, redirect_link, http.StatusMovedPermanently)
			}

			claims, err = VerifyAccessToken(newAccessToken.AccessToken, a.SecretKey)
			if err != nil {
				a.Lg.Info("Invalid or expired token")
				http.Redirect(w, r, redirect_link, http.StatusMovedPermanently)
				return
			}

			SetTokenCookie(
				AccessTokenCookieName,
				newAccessToken.AccessToken,
				time.Now().Add(15*time.Minute), w)
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
