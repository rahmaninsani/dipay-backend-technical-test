package usecase

import "github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/web"

type EmployeeUseCase interface {
	Create(payload web.EmployeeCreateRequest) (web.EmployeeCreateResponse, error)
}
