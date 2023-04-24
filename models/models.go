package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName     *string            `json:"first_name" validate:"required,min=2,max=30"`
	LastName      *string            `json:"last_name"  validate:"required,min=2,max=30"`
	Password      *string            `json:"password"   validate:"required,min=6"`
	Email         *string            `json:"email"      validate:"email,required"`
	Phone         *string            `json:"phone"      validate:"required"`
	Token         *string            `json:"token"`
	Refresh_Token *string            `json:"refresh_token"`
	Created_At    time.Time          `json:"created_at"`
	Updated_At    time.Time          `json:"updtaed_at"`
	User_ID       string             `json:"user_id"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        uint64             `json:"price" bson:"price"`
	Rating       *uint8             `json:"rating"`
	Image        *string            `json:"image"`
}

type Comment struct {
	ID        primitive.ObjectID
	UserID    primitive.ObjectID `bson:"_id"`
	ProductID primitive.ObjectID `bson:"_id"`
	Comment   string             `json:"comment"`
}
