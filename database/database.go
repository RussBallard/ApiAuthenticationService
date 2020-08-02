package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	mongodbUrl          = "mongodb://127.0.0.1:27017"
	mongoDefaultTimeout = 5 * time.Second
	err                 error
)

type BaseConnection struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func (mongoDB *BaseConnection) InitMongoDB() error {
	mongoDB.Client, err = mongo.NewClient(options.Client().ApplyURI(mongodbUrl))
	if err != nil {
		log.Fatal(err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeout)
	defer cancel()

	if err = mongoDB.Client.Connect(ctx); err != nil {
		return err
	}

	log.Println("Connected to MongoDB!")

	mongoDB.DB = mongoDB.Client.Database("authentication")
	return nil
}

func (mongoDB *BaseConnection) MongoCloseConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeout)
	defer cancel()

	if err = mongoDB.Client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed!")
}

func (mongoDB *BaseConnection) MongoSaveObject(databaseObject UserTokens) (*mongo.InsertOneResult, error) {
	collection := mongoDB.DB.Collection("user_tokens")

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeout)
	defer cancel()

	collectionResult, err := collection.InsertOne(ctx, databaseObject)
	return collectionResult, err
}

func (mongoDB *BaseConnection) MongoUpdateObject(databaseObject UserTokens) (*mongo.UpdateResult, error) {
	collection := mongoDB.DB.Collection("user_tokens")

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeout)
	defer cancel()

	collectionResult, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": databaseObject.GUID},
		bson.D{
			{"$set", bson.D{{"refresh_tokens", databaseObject.RefreshTokens}}},
		})
	return collectionResult, err
}
