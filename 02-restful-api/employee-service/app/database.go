package app

import (
	"context"
	"fmt"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func NewMongoDB() (context.Context, *mongo.Client) {
	constant := config.Constant
	
	var dbUri string
	if constant.DBUser == "" || constant.DBPassword == "" {
		dbUri = fmt.Sprintf("mongodb://%s:%s", constant.DBHost, constant.DBPort)
	} else {
		dbUri = fmt.Sprintf("mongodb://%s:%s@%s:%s", constant.DBUser,
			constant.DBPassword, constant.DBHost, constant.DBPort)
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Println("Failed to connect to the database: ", err.Error())
		panic(err)
	}
	
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Println("Failed to connect to the database: ", err.Error())
		panic(err)
	}
	
	log.Println("🚀 Connected successfully to the database")
	
	return ctx, client
}
