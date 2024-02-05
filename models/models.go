package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataProvider struct {
	OID       primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"firstname"`
	Email     string             `bson:"email"`
	Country   string             `bson:"country"`
	WebSite   string             `bson:"website"`
	DataTypes string             `bson:"data_types"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}
