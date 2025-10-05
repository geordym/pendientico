package middlewares

import (
	"context"
	"strings"

	"github.com/geordym/pendientico/infraestructure/configuration/security"
	"github.com/labstack/echo/v4"
)

type ClaimsKey struct{}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
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

		ctx := context.WithValue(c.Request().Context(), ClaimsKey{}, claims)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
