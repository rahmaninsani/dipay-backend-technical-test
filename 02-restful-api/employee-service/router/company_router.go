package router

import (
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/handler"
)

func NewCompanyRouter(group *echo.Group, companyHandler handler.CompanyHandler, employeeHandler handler.EmployeeHandler,
	middlewares []echo.MiddlewareFunc) {
	company := group.Group("/companies")
	
	company.POST("", companyHandler.Create, middlewares...)
	company.GET("", companyHandler.FindAll)
	company.PUT("/:id/set_active", companyHandler.SetActive, middlewares...)
	
	company.POST("/:company_id/employees", employeeHandler.Create, middlewares...)
}
