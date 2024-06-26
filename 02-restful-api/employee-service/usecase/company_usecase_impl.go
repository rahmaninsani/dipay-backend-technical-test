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

type CompanyUseCaseImpl struct {
	CompanyRepository repository.CompanyRepository
}

func NewCompanyUseCase(companyRepository repository.CompanyRepository) CompanyUseCase {
	return &CompanyUseCaseImpl{
		CompanyRepository: companyRepository,
	}
}

func (useCase CompanyUseCaseImpl) Create(payload web.CompanyCreateRequest) (web.CompanyCreateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	company := domain.Company{
		CompanyName: payload.CompanyName,
	}
	
	if payload.TelephoneNumber != nil {
		company.TelephoneNumber = *payload.TelephoneNumber
	}
	
	if payload.Address != nil {
		company.Address = *payload.Address
	}
	
	existedCompany, err := useCase.CompanyRepository.FindOne(ctx, company)
	if err != nil && err != echo.ErrNotFound {
		return web.CompanyCreateResponse{}, err
	}
	
	if existedCompany.ID != primitive.NilObjectID {
		return web.CompanyCreateResponse{}, echo.NewHTTPError(http.StatusConflict, "Company Name already exist")
	}
	
	company, err = useCase.CompanyRepository.Save(ctx, company)
	if err != nil {
		return web.CompanyCreateResponse{}, err
	}
	
	return helper.ToCompanyCreateResponse(company), nil
}

func (useCase CompanyUseCaseImpl) Update(payload web.CompanyUpdateRequest) (web.CompanyUpdateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	objectID, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		return web.CompanyUpdateResponse{}, echo.NewHTTPError(http.StatusUnprocessableEntity, "Data is not found")
	}
	
	company := domain.Company{
		ID: objectID,
	}
	
	if payload.CompanyName != nil {
		company.CompanyName = *payload.CompanyName
	}
	
	if payload.TelephoneNumber != nil {
		company.TelephoneNumber = *payload.TelephoneNumber
	}
	
	if payload.IsActive != nil {
		company.IsActive = *payload.IsActive
	}
	
	if payload.Address != nil {
		company.Address = *payload.Address
	}
	
	existedCompany, err := useCase.CompanyRepository.FindOne(ctx, company)
	if err != nil {
		if err == echo.ErrNotFound {
			return web.CompanyUpdateResponse{}, echo.NewHTTPError(http.StatusUnprocessableEntity, "Data is not found")
		}
		return web.CompanyUpdateResponse{}, err
	}
	
	if existedCompany.IsActive {
		return web.CompanyUpdateResponse{}, echo.NewHTTPError(http.StatusBadRequest, "Company is already active")
	}
	
	company, err = useCase.CompanyRepository.Update(ctx, company)
	if err != nil {
		return web.CompanyUpdateResponse{}, err
	}
	
	return helper.ToCompanyUpdateResponse(company), nil
}

func (useCase CompanyUseCaseImpl) FindAll() ([]web.CompanyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	companies, err := useCase.CompanyRepository.FindAll(ctx)
	if err != nil {
		return []web.CompanyResponse{}, err
	}
	
	if len(companies) == 0 {
		return []web.CompanyResponse{}, echo.NewHTTPError(http.StatusUnprocessableEntity, "Data is not found")
	}
	
	return helper.ToCompanyResponses(companies), nil
}
