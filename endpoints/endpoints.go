package endpoints

import (
	"ApiAuthenticationService/database"
	"ApiAuthenticationService/utils"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type BaseConnection struct {
	DatabaseData *database.BaseConnection
}

func (mongoDB *BaseConnection) CreatePairTokens(w http.ResponseWriter, r *http.Request) {
	var userData utils.UserGUID

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pairTokens := utils.GenerateJWT(userData.GUID)
	collection := mongoDB.DatabaseData.DB.Collection("user_tokens")
	guid := collection.FindOne(context.Background(), bson.M{"_id": userData.GUID})
	// TODO: Ограничить максимальное количество токенов
	if guid.Err() != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusConflict)
		return
	}

	databaseDocument := pairTokens.CreateDatabaseDocument(userData.GUID)
	collectionResult, err := mongoDB.DatabaseData.MongoSaveObject(databaseDocument)
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
	var userRefreshToken utils.UserRefreshToken
	var result database.UserTokens

	if err := json.NewDecoder(r.Body).Decode(&userRefreshToken); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	collection := mongoDB.DatabaseData.DB.Collection("user_tokens")
	if err := collection.FindOne(context.Background(), bson.M{"_id": userRefreshToken.GUID}).Decode(&result); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var isEqualToken = false
	var refreshTokenIndex int
	var refreshTokenData database.TokenType
	for index, refreshToken := range result.RefreshTokens {
		err := bcrypt.CompareHashAndPassword([]byte(refreshToken.Token), []byte(userRefreshToken.RefreshToken))
		if err == nil {
			isEqualToken = true
			refreshTokenIndex = index
			refreshTokenData = refreshToken
			break
		}
	}

	// Если токен не найден по хэшу в бд или истек его срок действия, то возвращаем 401 статус
	if isEqualToken == false || refreshTokenData.ExpiredAt.Unix() < time.Now().Unix() {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	pairTokens := utils.GenerateJWT(userRefreshToken.GUID)
	databaseDocument := pairTokens.CreateDatabaseDocument(userRefreshToken.GUID)

	// Производим замену найденного документа по индексу на полученный databaseDocument
	result.RefreshTokens[refreshTokenIndex] = databaseDocument.RefreshTokens[0]
	if _, err := mongoDB.DatabaseData.MongoUpdateObject(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Updated a single document: ", result.GUID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pairTokens)
}

func (mongoDB BaseConnection) DeleteOneToken(w http.ResponseWriter, r *http.Request) {

}

func (mongoDB BaseConnection) DeleteAllTokens(w http.ResponseWriter, r *http.Request) {

}
