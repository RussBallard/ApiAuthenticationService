package endpoints

import (
	"ApiAuthenticationService/database"
	"ApiAuthenticationService/utils"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type BaseConnection struct {
	DatabaseData *database.BaseConnection
}

func (mongoDB *BaseConnection) CreatePairTokens(w http.ResponseWriter, r *http.Request) {
	var rBody utils.UserGuid

	if err := json.NewDecoder(r.Body).Decode(&rBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pairTokens := utils.GenerateJWT(rBody.GUID)
	collection := mongoDB.DatabaseData.DB.Collection("user_tokens")
	guid := collection.FindOne(context.Background(), bson.M{"_id": rBody.GUID})
	// TODO: Ограничить максимальное количество токенов
	if guid.Err() != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusConflict)
		return
	}

	databaseDocument := pairTokens.CreateDatabaseDocument(rBody.GUID)
	collectionResult, err := mongoDB.DatabaseData.MongoSaveDocument(databaseDocument)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Inserted a single document: ", collectionResult.InsertedID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pairTokens)
}

func (mongoDB *BaseConnection) UpdatePairTokens(w http.ResponseWriter, r *http.Request) {
	var rBody utils.UserToken

	if err := json.NewDecoder(r.Body).Decode(&rBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbResult, err := mongoDB.DatabaseData.FindUserDocument(rBody.GUID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenIndex, err := rBody.FindEqualToken(dbResult)
	if err != nil {
		log.Printf("Token: %s ; Error: %s\n", err, rBody.RefreshToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	pairTokens := utils.GenerateJWT(rBody.GUID)
	databaseDocument := pairTokens.CreateDatabaseDocument(rBody.GUID)

	// Производим замену найденного документа по индексу на полученный databaseDocument
	dbResult.RefreshTokens[tokenIndex] = databaseDocument.RefreshTokens[0]
	if _, err := mongoDB.DatabaseData.MongoUpdateDocument(dbResult); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Updated a single document: ", dbResult.GUID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pairTokens)
}

func (mongoDB BaseConnection) DeleteOneToken(w http.ResponseWriter, r *http.Request) {
	var rBody utils.UserToken

	if err := json.NewDecoder(r.Body).Decode(&rBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbResult, err := mongoDB.DatabaseData.FindUserDocument(rBody.GUID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenIndex, err := rBody.FindEqualToken(dbResult)
	if err != nil {
		log.Printf("Token: %s ; Error: %s\n", err, rBody.RefreshToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	dbResult.RefreshTokens = append(dbResult.RefreshTokens[:tokenIndex], dbResult.RefreshTokens[tokenIndex+1:]...)
	if _, err := mongoDB.DatabaseData.MongoUpdateDocument(dbResult); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Removed single token: ", rBody.RefreshToken)
	w.WriteHeader(http.StatusOK)
}

func (mongoDB BaseConnection) DeleteAllTokens(w http.ResponseWriter, r *http.Request) {
	var rBody utils.UserToken

	if err := json.NewDecoder(r.Body).Decode(&rBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbResult, err := mongoDB.DatabaseData.FindUserDocument(rBody.GUID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if _, err = rBody.FindEqualToken(dbResult); err != nil {
		log.Printf("Token: %s ; Error: %s\n", err, rBody.RefreshToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	dbResult.RefreshTokens = nil
	if _, err := mongoDB.DatabaseData.MongoUpdateDocument(dbResult); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Removed all tokens from guid: ", dbResult.GUID)
	w.WriteHeader(http.StatusOK)
}
