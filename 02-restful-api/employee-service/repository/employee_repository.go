package repository

import (
	"context"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/domain"
)

type EmployeeRepository interface {
	Save(ctx context.Context, employee domain.Employee) (domain.Employee, error)
	FindOne(ctx context.Context, employee domain.Employee) (domain.Employee, error)
}
