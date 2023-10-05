package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	CompanyName     string             `bson:"company_name"`
	TelephoneNumber string             `bson:"telephone_number"`
	IsActive        bool               `bson:"is_active"`
	Address         string             `bson:"address"`
}
