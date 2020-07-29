package database

import (
	"ApiAuthenticationService/helpers"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func MongoConnection() *mongo.Client {
	clientOptions := options.Client().ApplyURI(helpers.MongodbUrl)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	return client
}

func (userTokens UserTokens) CreateTokensForUser(db *mongo.Client) {

}

func (userTokens UserTokens) UpdateTokensForUser(db *mongo.Client) {

}

func (userTokens UserTokens) DeleteOneTokenForUser(db *mongo.Client) {

}

func (userTokens UserTokens) DeleteAllTokensForUser(db *mongo.Client) {

}

func MongoCloseConnection(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed!")
}
