package mongoutil

import (
	"context"
	"errors"
	"net/url"

	"bitbucket.org/amaltheafs/pkg/logutil"
	"bitbucket.org/amaltheafs/pkg/sharedconfigs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cli *mongo.Client = nil

func init() {
	fUri = ""
}

var fUri string

func SetUriFromConfig(mongoDBConfig sharedconfigs.IMongoDBConfiguration) {
	fUri = mongoDBConfig.GetMongoDBURL()
}

// / Remove the password from the URL's User field
// / A better practice would be to use an environment variable without the password
func removePasswordFromMongoURI(uri string) (string, error) {
	parsedURL, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	// Remove the password from the URL's User field
	parsedURL.User = url.User(parsedURL.User.Username())

	return parsedURL.User.Username() + "@" + parsedURL.Host + parsedURL.Path, nil
}

func SetUriFromEnv() {
	SetUri("mongodb://localhost:27018")

	noPassURI, err := removePasswordFromMongoURI(fUri)
	if err != nil {
		logutil.Logger().Infof("Cannot log Mongodb URI")
	} else {
		logutil.Logger().Infof("Mongodb URI:%s", noPassURI)
	}
}

func SetUri(uri string) {
	fUri = uri
}

func GetUri() string {
	return fUri
}

func GetConnection(c context.Context) (*mongo.Client, error) {
	//TODO => handle context. Handle a pool of connection.
	SetUriFromEnv()
	if cli == nil {
		if fUri == "" {
			logutil.Logger().Errorf("Cannot connect to MongoDb, connection string is null")
			return nil, errors.New("Cannot connect to MongoDb, connection string is null")
		}
		var err error
		cli, err = mongo.Connect(c, options.Client().ApplyURI(fUri))
		if err != nil {
			return nil, errors.New("cannot connect to MongoDb")
		}
	}
	return cli, nil
}
func CloseConnection(ctx context.Context) error {
	return cli.Disconnect(ctx)
}
