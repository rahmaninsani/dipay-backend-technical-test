package repository

import (
	"context"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/domain"
)

type CompanyRepository interface {
	FindOne(ctx context.Context, company domain.Company) (domain.Company, error)
	FindAll(ctx context.Context) ([]domain.Company, error)
	Save(ctx context.Context, company domain.Company) (domain.Company, error)
	Update(ctx context.Context, company domain.Company) (domain.Company, error)
}
