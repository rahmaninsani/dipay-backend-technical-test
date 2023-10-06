package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/config"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepositoryImpl struct {
	Client *mongo.Client
}

func NewAdminRepository(client *mongo.Client) AdminRepository {
	return &AdminRepositoryImpl{
		Client: client,
	}
}

func (repository AdminRepositoryImpl) FindOne(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
	filter := bson.M{
		"username": admin.Username,
	}
	
	if err := repository.Client.
		Database(config.Constant.DBName).
		Collection("admins").
		FindOne(ctx, filter).Decode(&admin);
		err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Admin{}, fmt.Errorf("record not found")
		}
		return domain.Admin{}, err
	}
	
	return admin, nil
}
