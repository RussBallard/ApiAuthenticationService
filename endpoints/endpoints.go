package endpoints

import (
	"ApiAuthenticationService/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var CreateTokens = func(db *mongo.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res = helpers.SetResponseHeaders(res)
	}
}

var UpdateTokens = func(db *mongo.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res = helpers.SetResponseHeaders(res)
	}
}

var DeleteToken = func(db *mongo.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res = helpers.SetResponseHeaders(res)
	}
}

var DeleteTokens = func(db *mongo.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res = helpers.SetResponseHeaders(res)
	}
}
