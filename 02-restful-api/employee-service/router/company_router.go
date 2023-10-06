package router

import (
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/handler"
)

func NewCompanyRouter(group *echo.Group, companyHandler handler.CompanyHandler, middlewares []echo.MiddlewareFunc) {
	company := group.Group("/companies")
	
	company.POST("", companyHandler.Create, middlewares...)
}
