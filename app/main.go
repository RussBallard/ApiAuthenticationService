package main

import (
	"ApiAuthenticationService/handler"
	"ApiAuthenticationService/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	handler, err := handler.NewBaseHandler()
	if err != nil {
		log.Fatal(err)
		return
	}

	router := mux.NewRouter()

	router.HandleFunc(utils.CreateTokenURI, handler.CreatePairTokens).Methods(utils.MethodPOST)
	router.HandleFunc(utils.UpdateTokenURI, handler.UpdatePairTokens).Methods(utils.MethodPUT)
	router.HandleFunc(utils.DeleteTokensURI, handler.DeleteOneToken).Methods(utils.MethodDELETE)
	router.HandleFunc(utils.DeleteTokenURI, handler.DeleteAllTokens).Methods(utils.MethodDELETE)

	if err := http.ListenAndServe(utils.ApiPort, router); err != nil {
		log.Fatalf("ListenAndServe(): %s", err)
	}
}
