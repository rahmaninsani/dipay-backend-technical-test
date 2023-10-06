package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/config"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/helper"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/domain"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/web"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AdminUseCaseImpl struct {
	AdminRepository repository.AdminRepository
}

func NewAdminUseCase(adminRepository repository.AdminRepository) AdminUseCase {
	return &AdminUseCaseImpl{
		AdminRepository: adminRepository,
	}
}

func (useCase AdminUseCaseImpl) Login(payload web.AdminLoginRequest) (web.AdminLoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	admin, err := useCase.AdminRepository.FindOne(ctx, domain.Admin{
		Username: payload.Username,
	})
	if err != nil {
		if err.Error() == "record not found" {
			return web.AdminLoginResponse{}, fmt.Errorf("wrong username")
		}
		return web.AdminLoginResponse{}, err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(payload.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return web.AdminLoginResponse{}, fmt.Errorf("wrong username")
		}
	}
	
	accessTokenExpiresIn := time.Duration(config.Constant.AccessTokenExpiresIn) * time.Minute
	accessTokenSecretKey := config.Constant.AccessTokenSecretKey
	token, err := helper.GenerateToken(&admin, accessTokenExpiresIn, accessTokenSecretKey)
	if err != nil {
		return web.AdminLoginResponse{}, err
	}
	
	return helper.ToAdminLoginResponse(token), nil
}
