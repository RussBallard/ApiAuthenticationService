package main

import (
	"ApiAuthenticationService/database"
	"ApiAuthenticationService/endpoints"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	db := database.MongoConnection()
	defer database.MongoCloseConnection(db)

	router := mux.NewRouter()

	router.HandleFunc("/api/create/tokens", endpoints.CreateTokens(db)).Methods("POST")
	router.HandleFunc("/api/update/tokens", endpoints.UpdateTokens(db)).Methods("PUT")
	router.HandleFunc("/api/delete/token", endpoints.DeleteToken(db)).Methods("DELETE")
	router.HandleFunc("/api/delete/tokens", endpoints.DeleteTokens(db)).Methods("DELETE")

	server := &http.Server{Addr: ":8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown(): %s", err)
	}
}
