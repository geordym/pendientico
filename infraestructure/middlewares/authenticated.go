package middlewares

import (
    "strings"
    "github.com/labstack/echo/v4"
	"github.com/geordym/pendientico/configuration/security"

)


func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		idToken, err := security.Verifier.Verify(c.Request().Context(), tokenStr)
		if err != nil {
			return echo.ErrUnauthorized
		}

		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err != nil {
			return echo.ErrUnauthorized
		}

		c.Set("claims", claims)

		return next(c)
	}
}
