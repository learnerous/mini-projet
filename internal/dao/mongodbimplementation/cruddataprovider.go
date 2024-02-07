package mongodbimplementation

import (
	"context"
	"fmt"
	"log"
	"mini_projet/models"
	"mini_projet/pkg/core/errors"
	"mini_projet/pkg/logutil"
	"mini_projet/pkg/mongoutil"
	"mini_projet/pkg/rest/errorcodes/generic"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Collection
var ctx context.Context

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27018")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the database and collection variables
	db = client.Database("dataProvidersmanagement").Collection("dataproviders")
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
}

type MongoDBCRUDDataProvider struct {
}

const QUERY_TIMEOUT time.Duration = 5 * time.Second
const UNIQUE_FIELD_ERROR_CODE = "E11000"

func (mongoDBCRUDDataProvider *MongoDBCRUDDataProvider) CreateIndexes(ctx context.Context) error {
	client, err := mongoutil.GetConnection(ctx)
	if err != nil {
		logutil.Logger().Error(err.Error())
		return err
	}
	db := client.Database("dataProvidersmanagement")
	dataProviderCollection := db.Collection("dataproviers")
	if CreateIndexes(ctx, dataProviderCollection, []string{"clientCredentialApp"}) != nil {
		//TODO : uncomment this line when the index creation is fixed
		//return fmt.Errorf("error during index creation")
		return nil
	}
	return nil
}
func CreateIndexes(ctx context.Context, collection *mongo.Collection, fields []string) error {
	for _, field := range fields {
		indexModel := mongo.IndexModel{
			Keys:    bson.D{{field, 1}},
			Options: options.Index().SetUnique(true),
		}
		index, err := collection.Indexes().CreateOne(ctx, indexModel)
		if err != nil {
			return err
		}
		logutil.Logger().Infof("Index %s created on %s", index, collection.Name())
	}
	return nil
}
func (mongoDBCRUDDataProvider *MongoDBCRUDDataProvider) collection(ctx context.Context) *mongo.Collection {
	client, err := mongoutil.GetConnection(ctx)
	if err != nil {
		logutil.Logger().Error("Cannot get MongoDb connection:", err)
	}
	db := client.Database("dataProvidersmanagement")
	return db.Collection("dataproviders")
}

func (mongoDBCRUDDataProvider *MongoDBCRUDDataProvider) Create(ctx context.Context, dataprovider *models.DataProvider) (*models.DataProvider, error) {

	dataprovidercollection := mongoDBCRUDDataProvider.collection(ctx)

	contextWithCancel, cancel := context.WithTimeout(ctx, QUERY_TIMEOUT)
	defer cancel()

	if dataprovider.OID.IsZero() {
		dataprovider.OID = primitive.NewObjectID()
	}
	fmt.Println("dataprovider", dataprovider)
	dataprovider.CreatedAt = time.Now()
	dataprovider.UpdatedAt = time.Now()
	insertResult, err := dataprovidercollection.InsertOne(contextWithCancel, dataprovider)
	if err != nil {
		if strings.Contains(err.Error(), UNIQUE_FIELD_ERROR_CODE) {
			return nil, &errors.CustomError{
				Type: errors.DB,
				Code: generic.DB_UNIQUE.String(),
				Err:  err,
			}
		}
		return nil,
			&errors.CustomError{Type: errors.DB, Code: generic.UNKNOWN.String(), Err: err}
	}

	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		dataprovider.OID = oid
		return dataprovider, nil
	}
	return nil,
		&errors.CustomError{
			Type:    errors.DB,
			Message: "parsing/casting _id field from inserted doc",
			Code:    generic.UNKNOWN.String(),
		}
}
func (mongoDBCRUDDataProvider *MongoDBCRUDDataProvider) GetDataProvider(ctx context.Context) ([]*models.DataProvider, error) {
	dataprovidersCollection := mongoDBCRUDDataProvider.collection(ctx)
	contextWithCancel, cancel := context.WithTimeout(ctx, QUERY_TIMEOUT)
	defer cancel()
	filter := bson.M{}

	var dataProviders []*models.DataProvider
	cursor, err := dataprovidersCollection.Find(contextWithCancel, filter)
	if err != nil {
		return nil, &errors.CustomError{Type: errors.DB, Code: generic.UNKNOWN.String(), Err: err}
	}
	if err = cursor.All(contextWithCancel, &dataProviders); err != nil {
		return nil, &errors.CustomError{Type: errors.DB, Code: generic.UNKNOWN.String(), Err: err}
	}

	return dataProviders, nil
}

func (mongoDBCRUDDataProvider *MongoDBCRUDDataProvider) GetDataProviderByID(ctx context.Context, id string) (*models.DataProvider, error) {
	dataproviderId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, &errors.CustomError{Type: errors.INTERNALS, Code: generic.INVALID_OID.String(), Err: err}
	}
	dataproviderCollection := mongoDBCRUDDataProvider.collection(ctx)
	contextWithCancel, cancel := context.WithTimeout(ctx, QUERY_TIMEOUT)
	defer cancel()
	cursor, err := dataproviderCollection.Find(ctx, bson.M{"_id": dataproviderId})
	if err != nil {
		return nil, &errors.CustomError{Type: errors.DB, Code: generic.UNKNOWN.String(), Err: err}
	}
	var result []*models.DataProvider
	if err = cursor.All(contextWithCancel, &result); err != nil {
		return nil, &errors.CustomError{Type: errors.DB, Code: generic.UNKNOWN.String(), Err: err}
	}
	if len(result) == 0 {
		return nil, nil
	}
	return result[0], nil
}
func (mongoDBCRUDDataProvider *MongoDBCRUDDataProvider) UpdateDataProvider(ctx context.Context, id string, newDataprovider *models.DataProvider) (*models.DataProvider, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	dataproviderCollection := mongoDBCRUDDataProvider.collection(ctx)
	contextWithCancel, cancel := context.WithTimeout(ctx, QUERY_TIMEOUT)
	defer cancel()
	//updatedDataprovider := UpdatedDataprovider(newDataprovider)
	var updatedDataprovider = &models.DataProvider{}
	updatedFields := bson.M{}
	if newDataprovider.Name != "" {
		updatedFields["name"] = newDataprovider.Name
	}
	if newDataprovider.Email != "" {
		updatedFields["email"] = newDataprovider.Email
	}
	if newDataprovider.Country != "" {
		updatedFields["country"] = newDataprovider.Country
	}
	if newDataprovider.WebSite != "" {
		updatedFields["website"] = newDataprovider.WebSite
	}
	if newDataprovider.DataTypes != "" {
		updatedFields["data_types"] = newDataprovider.DataTypes
	}
	updatedFields["updated_at"] = time.Now()
	err = dataproviderCollection.FindOneAndUpdate(
		contextWithCancel,
		bson.M{"_id": oid},
		bson.M{"$set": updatedFields},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(updatedDataprovider)
	if err != nil {
		return nil,
			&errors.CustomError{Type: errors.DB, Code: generic.UNKNOWN.String(), Err: err}

	}
	return updatedDataprovider, nil

}

func (mongoDBCRUDDataProvider *MongoDBCRUDDataProvider) DeleteDataProvider(ctx context.Context, id string) (bool, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false,
			&errors.CustomError{Type: errors.INTERNALS, Code: generic.INVALID_OID.String(), Err: err}
	}

	usersCollection := mongoDBCRUDDataProvider.collection(ctx)
	contextWithCancel, cancel := context.WithTimeout(ctx, QUERY_TIMEOUT)
	defer cancel()

	deleteResult, err := usersCollection.DeleteOne(contextWithCancel, bson.M{"_id": oid})
	if err != nil {
		return deleteResult.DeletedCount == 1, &errors.CustomError{Type: errors.DB, Code: generic.UNKNOWN.String(), Err: err}
	}

	return deleteResult.DeletedCount == 1, nil

}
