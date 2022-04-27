package models

import (
	"time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	FirstName  string             `json:"firstName" bson:"firstName"`
	SurName    string             `json:"surName" bson:"surName"`
	LastName   string             `json:"lastName" bson:"lastName"`
	FullName   string             `json:"fullName" bson:"fullName"`
	Email      string             `json:"email" bson:"email"`
	Phone      string             `json:"phone" bson:"phone" validate:"required"`
	UserType   string             `json:"userType" bson:"userType"`
	UserStatus string             `json:"userStatus" bson:"userStatus"`
    UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
}
