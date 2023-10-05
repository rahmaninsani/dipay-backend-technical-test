package helper

import (
	"context"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/config"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func InitCollection(client *mongo.Client, ctx context.Context) {
	constant := config.Constant
	
	adminCollection := client.Database(constant.DBName).Collection("admins")
	count, err := adminCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Println("Failed to count admin documents: ", err.Error())
		panic(err)
	}
	
	if count > 0 {
		log.Println("Admin data already exists")
		return
	}
	
	admin := domain.Admin{
		Username: "admin",
		Password: "admin", // TODO: Hash password
	}
	
	_, err = client.Database(constant.DBName).Collection("admins").InsertOne(ctx, admin)
	if err != nil {
		log.Println("Failed to insert admin data: ", err.Error())
		panic(err)
	}
	
	log.Println("Admin data inserted successfully")
}
