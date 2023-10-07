package usecase

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/helper"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/domain"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/web"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type EmployeeUseCaseImpl struct {
	EmployeeRepository repository.EmployeeRepository
	CompanyRepository  repository.CompanyRepository
}

func NewEmployeeUseCase(employeeRepository repository.EmployeeRepository, companyRepository repository.CompanyRepository) EmployeeUseCase {
	return &EmployeeUseCaseImpl{
		EmployeeRepository: employeeRepository,
		CompanyRepository:  companyRepository,
	}
}

func (useCase EmployeeUseCaseImpl) Create(payload web.EmployeeCreateRequest) (web.EmployeeCreateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	objectID, err := primitive.ObjectIDFromHex(payload.CompanyID)
	if err != nil {
		return web.EmployeeCreateResponse{}, echo.NewHTTPError(http.StatusBadRequest, "Invalid Company ID")
	}
	
	existedCompany, err := useCase.CompanyRepository.FindOne(ctx, domain.Company{ID: objectID})
	if err != nil {
		if err == echo.ErrNotFound {
			return web.EmployeeCreateResponse{}, echo.NewHTTPError(http.StatusUnprocessableEntity, "Company is not found")
		}
		return web.EmployeeCreateResponse{}, err
	}
	
	employee := domain.Employee{
		Name:      payload.Name,
		Email:     payload.Email,
		JobTitle:  payload.JobTitle,
		CompanyID: existedCompany.ID,
	}
	
	if payload.PhoneNumber != nil {
		employee.PhoneNumber = *payload.PhoneNumber
	}
	
	existedEmployee, err := useCase.EmployeeRepository.FindOne(ctx, employee)
	if err != nil && err != echo.ErrNotFound {
		return web.EmployeeCreateResponse{}, err
	}
	
	if existedEmployee.ID != primitive.NilObjectID {
		return web.EmployeeCreateResponse{}, echo.NewHTTPError(http.StatusConflict, "Email already exist")
	}
	
	employee, err = useCase.EmployeeRepository.Save(ctx, employee)
	if err != nil {
		return web.EmployeeCreateResponse{}, err
	}
	
	return helper.ToEmployeeCreateResponse(employee, existedCompany), nil
}
