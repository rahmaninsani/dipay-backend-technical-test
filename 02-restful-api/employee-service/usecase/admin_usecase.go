package usecase

import "github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/web"

type AdminUseCase interface {
	Login(payload web.AdminLoginRequest) (web.AdminLoginResponse, error)
}
