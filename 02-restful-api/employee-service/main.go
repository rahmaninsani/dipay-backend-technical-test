package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/app"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/config"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/exception"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/helper"
	"log"
)

func init() {
	err := config.LoadConstant()
	if err != nil {
		log.Fatalln("Failed to load environment variables\n", err.Error())
	}
}

func main() {
	client, ctx, cancel := app.NewMongoDB()
	defer app.CloseMongoDB(client, ctx, cancel)
	helper.InitCollection(client, ctx)
	
	e := echo.New()
	e.HTTPErrorHandler = exception.HTTPErrorHandler
	
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
		Output: e.Logger.Output(),
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	
	address := fmt.Sprintf(":%s", config.Constant.AppPort)
	e.Logger.Fatal(e.Start(address))
}
