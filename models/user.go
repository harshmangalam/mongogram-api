package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountDeactivate struct {
	From time.Time `bson:"from" json:"-"`
	To   time.Time `bson:"to" json:"-"`
}
type User struct {
	Id           primitive.ObjectID `bson:"_id" json:"_id"`
	Name         string             `bson:"name" json:"name"`
	Username     string             `bson:"username" json:"username"`
	Password     string             `bson:"password" json:"-"`
	Email        string             `bson:"email" json:"email"`
	Phone        string             `bson:"phone" json:"phone"`
	Birthday     time.Time          `bson:"birthday" json:"birthday"`
	Bio          string             `bson:"bio" json:"bio"`
	ResetPassOtp string             `bson:"resetPassOtp" json:"-"`
	Deactivate   *AccountDeactivate `bson:"deactivate" json:"-"`
	IsActive     bool               `bson:"isActive" json:"isActive"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
