package repository

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/config"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepositoryImpl struct {
	Client *mongo.Client
}

func NewEmployeeRepository(client *mongo.Client) EmployeeRepository {
	return &EmployeeRepositoryImpl{
		Client: client,
	}
}

func (repository EmployeeRepositoryImpl) Save(ctx context.Context, employee domain.Employee) (domain.Employee, error) {
	result, err := repository.Client.
		Database(config.Constant.DBName).
		Collection("employees").
		InsertOne(ctx, employee)
	if err != nil {
		return domain.Employee{}, err
	}
	
	employee.ID = result.InsertedID.(primitive.ObjectID)
	return employee, nil
}

func (repository EmployeeRepositoryImpl) FindOne(ctx context.Context, employee domain.Employee) (domain.Employee, error) {
	filter := bson.M{}
	
	if employee.ID != primitive.NilObjectID {
		filter["_id"] = employee.ID
	} else {
		filter["email"] = employee.Email
	}
	
	if err := repository.Client.
		Database(config.Constant.DBName).
		Collection("employees").
		FindOne(ctx, filter).
		Decode(&employee);
		err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Employee{}, echo.ErrNotFound
		}
		
		return domain.Employee{}, err
	}
	
	return employee, nil
}
