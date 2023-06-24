package middleware

import (
	"github.com/labstack/echo"
)

// SecretStaticAuthentication represent static secret key middleware for internal communication authentication
func SecretStaticAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ac := c.(*util.CustomApplicationContext)
		// conf := ac.SharedConf

		// tokenHeader := c.Request().Header.Get("Authorization")

		// // if token is empty return missing authentication token response
		// if tokenHeader == "" {
		// 	return ac.CustomResponse("failed", nil, "Missing Authenticastion Token", "[FAILED][FINANCE-SERVICE][MIDDLEWARE][JWTAuthenctication] Missing Authenticastion Token", http.StatusUnauthorized, nil)
		// }

		// if tokenHeader != conf.GetTokenSecretStatic() {
		// 	return ac.CustomResponse("failed", nil, "Invalid Token", fmt.Sprintf("[FAILED][FINANCE-SERVICE][MIDDLEWARE][JWTAuthenctication] invalid internal token"), http.StatusUnauthorized, nil)
		// }
		return next(c)
	}
}
