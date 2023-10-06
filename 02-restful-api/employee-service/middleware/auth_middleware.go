package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/config"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/helper"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/domain"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/repository"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware(adminRepository repository.AdminRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.Contains(authHeader, "Bearer") {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}
			
			tokenString := strings.SplitN(authHeader, " ", 2)[1]
			tokenClaims, err := helper.ValidateToken(tokenString, config.Constant.AccessTokenSecretKey)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}
			
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			
			username := tokenClaims["username"].(string)
			user, err := adminRepository.FindOne(ctx, domain.Admin{Username: username})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}
			
			c.Set("user", user)
			return next(c)
		}
	}
}
