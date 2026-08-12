package middleware

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"p2-ip-kartikajls/internal/helper"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// =========================================
		// Authorization Header
		// =========================================

		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return helper.ResponseError(
				c,
				http.StatusUnauthorized,
				"Authorization header required",
				nil,
			)
		}

		// Format: Bearer <token>
		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			return helper.ResponseError(
				c,
				http.StatusUnauthorized,
				"Invalid authorization header format",
				nil,
			)
		}

		tokenString := parts[1]

		// =========================================
		// JWT Secret
		// =========================================

		secret := os.Getenv("JWT_SECRET")

		if secret == "" {
			return helper.ResponseError(
				c,
				http.StatusInternalServerError,
				"JWT_SECRET belum diset",
				nil,
			)
		}

		// =========================================
		// Parse JWT
		// =========================================

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {

				if token.Method != jwt.SigningMethodHS256 {
					return nil, echo.NewHTTPError(
						http.StatusUnauthorized,
						"invalid signing method",
					)
				}

				return []byte(secret), nil
			},
		)

		if err != nil || !token.Valid {
			return helper.ResponseError(
				c,
				http.StatusUnauthorized,
				"Invalid or expired token",
				nil,
			)
		}

		// =========================================
		// Claims
		// =========================================

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			return helper.ResponseError(
				c,
				http.StatusUnauthorized,
				"Invalid token claims",
				nil,
			)
		}

		// =========================================
		// User ID
		// =========================================

		userIDValue, ok := claims["user_id"]

		if !ok {
			return helper.ResponseError(
				c,
				http.StatusUnauthorized,
				"User ID not found in token",
				nil,
			)
		}

		userIDString, ok := userIDValue.(string)

		if !ok {
			return helper.ResponseError(
				c,
				http.StatusUnauthorized,
				"Invalid user ID in token",
				nil,
			)
		}

		userID, err := strconv.ParseInt(
			userIDString,
			10,
			64,
		)

		if err != nil {
			return helper.ResponseError(
				c,
				http.StatusUnauthorized,
				"Invalid user ID",
				nil,
			)
		}

		// =========================================
		// Email
		// =========================================

		email, _ := claims["email"].(string)

		// =========================================
		// Role
		// =========================================

		role, ok := claims["role"].(string)

		if !ok || role == "" {
			return helper.ResponseError(
				c,
				http.StatusForbidden,
				"User role not found",
				nil,
			)
		}

		// =========================================
		// Simpan ke Echo Context
		// =========================================

		c.Set("user_id", userID)
		c.Set("email", email)
		c.Set("role", role)

		return next(c)
	}
}
