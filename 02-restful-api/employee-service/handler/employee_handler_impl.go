package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/helper"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/web"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/usecase"
	
	"net/http"
)

type EmployeeHandlerImpl struct {
	EmployeeUseCase usecase.EmployeeUseCase
}

func NewEmployeeHandler(employeeUseCase usecase.EmployeeUseCase) EmployeeHandler {
	return &EmployeeHandlerImpl{
		EmployeeUseCase: employeeUseCase,
	}
}

func (handler EmployeeHandlerImpl) Create(c echo.Context) error {
	companyID := c.Param("company_id")
	payload := web.EmployeeCreateRequest{
		CompanyID: companyID,
	}
	
	if err := c.Bind(&payload); err != nil {
		return err
	}
	
	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	isValidJobTitle := payload.JobTitle.IsValid()
	if !isValidJobTitle {
		allowedJobTitles := payload.JobTitle.AllowedValues()
		message := fmt.Sprintf("%s is not one of the allowed values: %s", payload.JobTitle, allowedJobTitles)
		return echo.NewHTTPError(http.StatusBadRequest, message)
	}
	
	employeeCreateResponse, err := handler.EmployeeUseCase.Create(payload)
	if err != nil {
		return err
	}
	
	response := helper.ToResponse(http.StatusCreated, employeeCreateResponse, "Success")
	return c.JSON(http.StatusCreated, response)
}
