package main

import (
	"ApiAuthenticationService/database"
	"ApiAuthenticationService/endpoints"
	"ApiAuthenticationService/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	mongoDB := database.BaseConnection{}
	if err := mongoDB.InitMongoDB(); err != nil {
		log.Fatal(err)
		return
	}

	defer mongoDB.MongoCloseConnection()
	databaseConnection := endpoints.BaseConnection{DatabaseData: &mongoDB}

	router := mux.NewRouter()

	router.HandleFunc("/api/create/tokens", databaseConnection.CreatePairTokens).Methods("POST")
	router.HandleFunc("/api/update/tokens", databaseConnection.UpdatePairTokens).Methods("PUT")
	router.HandleFunc("/api/delete/token", utils.DeleteMethodHeaders(databaseConnection.DeleteOneToken)).Methods("DELETE")
	router.HandleFunc("/api/delete/tokens", utils.DeleteMethodHeaders(databaseConnection.DeleteAllTokens)).Methods("DELETE")

	if err := http.ListenAndServe(":"+utils.ServerPort, router); err != nil {
		log.Fatalf("ListenAndServe(): %s", err)
	}
}
