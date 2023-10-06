package usecase

import "github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/web"

type CompanyUseCase interface {
	Create(payload web.CompanyCreateRequest) (web.CompanyCreateResponse, error)
}
