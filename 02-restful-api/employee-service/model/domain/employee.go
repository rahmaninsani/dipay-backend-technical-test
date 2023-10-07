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

func (jt JobTitle) IsValid() bool {
	switch jt {
	case Manager, Director, Staff:
		return true
	}
	return false
}

func (jt JobTitle) AllowedValues() []string {
	return []string{
		Manager.String(),
		Director.String(),
		Staff.String(),
	}
}

type Employee struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Email       string             `bson:"email"`
	PhoneNumber string             `bson:"phone_number"`
	JobTitle    JobTitle           `bson:"jobtitle"`
	CompanyID   primitive.ObjectID `bson:"company_id"`
}
