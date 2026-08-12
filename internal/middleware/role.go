package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"p2-ip-kartikajls/internal/helper"
)

func RoleMiddleware(roles ...string) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			roleValue := c.Get("role")

			role, ok := roleValue.(string)

			if !ok || role == "" {
				return helper.ResponseError(
					c,
					http.StatusForbidden,
					"User role not found",
					nil,
				)
			}

			for _, allowedRole := range roles {

				if role == allowedRole {
					return next(c)
				}
			}

			return helper.ResponseError(
				c,
				http.StatusForbidden,
				"Access denied",
				nil,
			)
		}
	}
}
