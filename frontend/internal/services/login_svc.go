package services

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// )

// func postLogin(c context.Context) {

// 	//replace echo context
// 	username := c.FormValue("username")
// 	password := c.FormValue("password")
// 	loggedInUser, err := .LoginUser(username, password)
// 	if err != nil {
// 		fmt.Println(err)
// 		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid login.")
// 	}

// 	// call backend to get user
// 	err = services.GenerateTokensAndSetCookies(loggedInUser, c)

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusUnauthorized, "Token failed to be generated.")
// 	}

// 	return c.Redirect(http.StatusMovedPermanently, "/chat")
// }
