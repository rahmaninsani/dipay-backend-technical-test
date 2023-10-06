package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/helper"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/web"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/usecase"
	
	"net/http"
)

type AdminHandlerImpl struct {
	AdminUseCase usecase.AdminUseCase
}

func NewAdminHandler(adminUseCase usecase.AdminUseCase) AdminHandler {
	return &AdminHandlerImpl{
		AdminUseCase: adminUseCase,
	}
}

func (handler AdminHandlerImpl) Login(c echo.Context) error {
	var payload web.AdminLoginRequest
	if err := c.Bind(&payload); err != nil {
		return err
	}
	
	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	adminLoginResponse, err := handler.AdminUseCase.Login(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	response := helper.ToResponse(http.StatusCreated, adminLoginResponse, "Success")
	return c.JSON(http.StatusCreated, response)
}
