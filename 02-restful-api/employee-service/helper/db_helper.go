package helper

import (
	"context"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/config"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func InitCollection(client *mongo.Client) {
	constant := config.Constant
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
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
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to generate hashed password: ", err.Error())
		panic(err)
	}
	
	admin := domain.Admin{
		Username: "admin",
		Password: string(hashedPassword),
	}
	
	_, err = client.Database(constant.DBName).Collection("admins").InsertOne(ctx, admin)
	if err != nil {
		log.Println("Failed to insert admin data: ", err.Error())
		panic(err)
	}
	
	log.Println("Admin data inserted successfully")
}
