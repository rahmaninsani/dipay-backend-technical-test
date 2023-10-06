package router

import (
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/handler"
)

func NewAdminRouter(group *echo.Group, adminHandler handler.AdminHandler, middlewares []echo.MiddlewareFunc) {
	admin := group.Group("/admins")
	
	admin.POST("/login", adminHandler.Login)
}
