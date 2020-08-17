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
	mongodbUrl          = "mongodb://admin:password@mongo1,mongo2,mongo3/authentication?replicaSet=rs0&authSource=admin&w=1&"
	mongoDefaultTimeout = 5 * time.Second
)

type BaseConnection struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func (mongoDB *BaseConnection) InitMongoDB() error {
	mongoDB.Client, _ = mongo.NewClient(options.Client().ApplyURI(mongodbUrl))
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeout)
	defer cancel()

	err := mongoDB.Client.Connect(ctx)
	if err = mongoDB.Client.Ping(ctx, nil); err != nil {
		return err
	}

	log.Println("Connected to MongoDB!")

	mongoDB.DB = mongoDB.Client.Database("authentication")
	return nil
}

func (mongoDB *BaseConnection) MongoCloseConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeout)
	defer cancel()

	if err := mongoDB.Client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed!")
}

func (mongoDB *BaseConnection) FindUserDocument(guid string) (UserTokens, error) {
	var dbResult UserTokens
	collection := mongoDB.DB.Collection("user_tokens")
	err := collection.FindOne(context.Background(), bson.M{"_id": guid}).Decode(&dbResult)
	return dbResult, err
}

func (mongoDB *BaseConnection) MongoSaveDocument(databaseDocument UserTokens) (*mongo.InsertOneResult, error) {
	collection := mongoDB.DB.Collection("user_tokens")

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeout)
	defer cancel()

	collectionResult, err := collection.InsertOne(ctx, databaseDocument)
	return collectionResult, err
}

func (mongoDB *BaseConnection) MongoUpdateDocument(databaseDocument UserTokens) (*mongo.UpdateResult, error) {
	collection := mongoDB.DB.Collection("user_tokens")

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeout)
	defer cancel()

	collectionResult, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": databaseDocument.GUID},
		bson.D{{"$set", bson.D{{"refresh_tokens", databaseDocument.RefreshTokens}}}})
	return collectionResult, err
}
