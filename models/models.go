package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataProvider struct {
	OID       primitive.ObjectID `bson:"_id "`
	Name      string             `bson:"firstname"`
	Email     string             `bson:"email"`
	Country   string             `bson:"country"`
	WebSite   string             `bson:"website"`
	DataTypes string             `bson:"data_types"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
