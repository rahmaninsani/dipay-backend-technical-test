package helper

import (
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/domain"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/web"
	"strconv"
)

func ToResponse(code int, data any, message string) web.Response {
	return web.Response{
		Status:  code,
		Code:    strconv.Itoa(code),
		Data:    data,
		Message: message,
	}
}

func ToAdminLoginResponse(token string) web.AdminLoginResponse {
	return web.AdminLoginResponse{
		Token: token,
	}
}

func ToCompanyCreateResponse(company domain.Company) web.CompanyCreateResponse {
	return web.CompanyCreateResponse{
		ID: company.ID.Hex(),
	}
}

func ToCompanyResponse(company domain.Company) web.CompanyResponse {
	return web.CompanyResponse{
		ID:              company.ID.Hex(),
		CompanyName:     company.CompanyName,
		TelephoneNumber: company.TelephoneNumber,
		IsActive:        company.IsActive,
		Address:         company.Address,
	}
}

func ToCompanyResponses(companies []domain.Company) []web.CompanyResponse {
	var responses []web.CompanyResponse
	
	for _, company := range companies {
		responses = append(responses, ToCompanyResponse(company))
	}
	
	return responses
}

func ToCompanyUpdateResponse(company domain.Company) web.CompanyUpdateResponse {
	return web.CompanyUpdateResponse{
		ID:       company.ID.Hex(),
		IsActive: company.IsActive,
	}
}
