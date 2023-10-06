package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/helper"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/web"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/usecase"
	
	"net/http"
)

type CompanyHandlerImpl struct {
	CompanyUseCase usecase.CompanyUseCase
}

func NewCompanyHandler(companyUseCase usecase.CompanyUseCase) CompanyHandler {
	return &CompanyHandlerImpl{
		CompanyUseCase: companyUseCase,
	}
}

func (handler CompanyHandlerImpl) Create(c echo.Context) error {
	var payload web.CompanyCreateRequest
	if err := c.Bind(&payload); err != nil {
		return err
	}
	
	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	companyCreateResponse, err := handler.CompanyUseCase.Create(payload)
	if err != nil {
		return err
	}
	
	response := helper.ToResponse(http.StatusCreated, companyCreateResponse, "Success")
	return c.JSON(http.StatusCreated, response)
}

func (handler CompanyHandlerImpl) FindAll(c echo.Context) error {
	companyResponses, err := handler.CompanyUseCase.FindAll()
	if err != nil {
		return err
	}
	
	modifiedResponse := map[string]any{
		"count": len(companyResponses),
		"rows":  companyResponses,
	}
	
	response := helper.ToResponse(http.StatusOK, modifiedResponse, "Success")
	return c.JSON(http.StatusOK, response)
}
