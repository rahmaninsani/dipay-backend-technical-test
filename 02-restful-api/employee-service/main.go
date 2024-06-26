package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/app"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/config"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/exception"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/handler"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/helper"
	customMiddleware "github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/middleware"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/repository"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/router"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/usecase"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/validation"
	"log"
)

func init() {
	err := config.LoadConstant()
	if err != nil {
		log.Fatalln("Failed to load environment variables\n", err.Error())
	}
}

func main() {
	ctx, client := app.NewMongoDB()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	helper.InitCollection(client)
	
	e := echo.New()
	e.HTTPErrorHandler = exception.HTTPErrorHandler
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
		Output: e.Logger.Output(),
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	
	adminRepository := repository.NewAdminRepository(client)
	companyRepository := repository.NewCompanyRepository(client)
	employeeRepository := repository.NewEmployeeRepository(client)
	
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	companyUseCase := usecase.NewCompanyUseCase(companyRepository)
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepository, companyRepository)
	
	adminHandler := handler.NewAdminHandler(adminUseCase)
	companyHandler := handler.NewCompanyHandler(companyUseCase)
	employeeHandler := handler.NewEmployeeHandler(employeeUseCase)
	
	authMiddleware := customMiddleware.AuthMiddleware(adminRepository)
	
	api := e.Group("/api")
	router.NewAdminRouter(api, adminHandler, nil)
	router.NewCompanyRouter(api, companyHandler, employeeHandler, []echo.MiddlewareFunc{authMiddleware})
	
	address := fmt.Sprintf(":%s", config.Constant.AppPort)
	e.Logger.Fatal(e.Start(address))
}
