package repository

import (
	"context"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/domain"
)

type AdminRepository interface {
	FindOne(ctx context.Context, admin domain.Admin) (domain.Admin, error)
}
