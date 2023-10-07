package web

import "github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/domain"

type EmployeeCreateRequest struct {
	Name        string          `validate:"required,min=2,max=50" json:"name"`
	Email       string          `validate:"required,email,min=5,max=255" json:"email"`
	PhoneNumber *string         `validate:"omitempty,min=8,max=16" json:"phone_number"`
	JobTitle    domain.JobTitle `validate:"required,oneof=manager director staff" json:"jobtitle"`
	CompanyID   string          `validate:"required" json:"company_id"`
}
