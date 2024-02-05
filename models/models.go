package models

type DataProvider struct {
	OID       int    `bson:"_id"`
	Name      string `bson:"firstname"`
	Email     string `bson:"email"`
	Country   string `bson:"country"`
	CreatedAt string `bson:"created_at"`
	UpdatedAt string `bson:"updated_at"`
}