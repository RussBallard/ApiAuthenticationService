package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
)

var (
	SigningKey = os.Getenv("SIGNING_KEY")
	ServerPort = os.Getenv("SERVER_PORT")
	MongodbUrl = os.Getenv("MONGODB_URL")
)

func SetResponseHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Add("Content-Type", "application/json")
	return w
}

func GenerateJWT(body map[string]string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)

	for key, value := range body {
		claims[key] = value
	}

	tokenString, err := token.SignedString(SigningKey)

	if err != nil {
		log.Fatalf("ERORR: %s", err)
		return "", err
	}

	return tokenString, nil
}

func CompareHeaders(requestHeaders map[string]string) bool {
	HeadersJWT := map[string]string{
		"typ": "JWT",
		"alg": "HS512",
	}

	if len(requestHeaders) != len(HeadersJWT) {
		return false
	}

	for k, v := range requestHeaders {
		if HeadersJWT[k] != v {
			return false
		}
	}
	return true
}
