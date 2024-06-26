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

type CompanyRepositoryImpl struct {
	Client *mongo.Client
}

func NewCompanyRepository(client *mongo.Client) CompanyRepository {
	return &CompanyRepositoryImpl{
		Client: client,
	}
}

func (repository CompanyRepositoryImpl) FindOne(ctx context.Context, company domain.Company) (domain.Company, error) {
	filter := bson.M{}
	
	if company.ID != primitive.NilObjectID {
		filter["_id"] = company.ID
	} else {
		filter["company_name"] = company.CompanyName
	}
	
	if err := repository.Client.
		Database(config.Constant.DBName).
		Collection("companies").
		FindOne(ctx, filter).
		Decode(&company);
		err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Company{}, echo.ErrNotFound
		}
		
		return domain.Company{}, err
	}
	
	return company, nil
}

func (repository CompanyRepositoryImpl) FindAll(ctx context.Context) ([]domain.Company, error) {
	cursor, err := repository.Client.
		Database(config.Constant.DBName).
		Collection("companies").
		Find(ctx, bson.M{})
	if err != nil {
		return []domain.Company{}, err
	}
	
	var companies []domain.Company
	if err := cursor.All(ctx, &companies); err != nil {
		return []domain.Company{}, err
	}
	
	return companies, nil
}

func (repository CompanyRepositoryImpl) Save(ctx context.Context, company domain.Company) (domain.Company, error) {
	result, err := repository.Client.
		Database(config.Constant.DBName).
		Collection("companies").
		InsertOne(ctx, company)
	if err != nil {
		return domain.Company{}, err
	}
	
	company.ID = result.InsertedID.(primitive.ObjectID)
	return company, nil
}

func (repository CompanyRepositoryImpl) Update(ctx context.Context, company domain.Company) (domain.Company, error) {
	filter := bson.M{
		"_id": company.ID,
	}
	
	companyUpdate := bson.M{}
	if company.CompanyName != "" {
		companyUpdate["company_name"] = company.CompanyName
	}
	
	if company.TelephoneNumber != "" {
		companyUpdate["telephone_number"] = company.TelephoneNumber
	}
	
	if company.IsActive || !company.IsActive {
		companyUpdate["is_active"] = company.IsActive
	}
	
	if company.Address != "" {
		companyUpdate["address"] = company.Address
	}
	
	update := bson.M{
		"$set": companyUpdate,
	}
	
	if _, err := repository.Client.
		Database(config.Constant.DBName).
		Collection("companies").
		UpdateOne(ctx, filter, update);
		err != nil {
		return domain.Company{}, err
	}
	
	return company, nil
}
