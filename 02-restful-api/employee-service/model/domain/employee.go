package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobTitle string

const (
	Manager  JobTitle = "manager"
	Director JobTitle = "director"
	Staff    JobTitle = "staff"
)

func (jt JobTitle) String() string {
	return string(jt)
}

type Employee struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Email       string             `bson:"email"`
	PhoneNumber string             `bson:"phone_number"`
	JobTitle    JobTitle           `bson:"jobtitle"`
	CompanyID   primitive.ObjectID `bson:"_id,omitempty"`
}
