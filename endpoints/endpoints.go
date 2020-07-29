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

var DeleteOneToken = func(db *mongo.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res = helpers.SetResponseHeaders(res)
	}
}

var DeleteAllTokens = func(db *mongo.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res = helpers.SetResponseHeaders(res)
	}
}
