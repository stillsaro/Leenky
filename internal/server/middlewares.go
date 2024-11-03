package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) JwtAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("Jwt")
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "Missing or invalid token")
			}

			claims, err := s.Auth.JwtValidator(cookie.Value)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "invalid token")
			}
			c.Set("username", claims.Username)
			return next(c)
		}
	}
}
