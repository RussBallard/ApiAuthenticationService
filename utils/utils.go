package utils

import (
	"ApiAuthenticationService/database"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	SigningKey = []byte("my_secret_key")
	ServerPort = "8000"
)

type UserRefreshToken struct {
	GUID         string `json:"guid"`
	RefreshToken string `json:"refresh_token"`
}

type UserGUID struct {
	GUID string `json:"guid"`
}

type Claims struct {
	GUID string `json:"guid"`
	jwt.StandardClaims
}

type Tokens struct {
	AccessToken  TokenType `json:"access_token"`
	RefreshToken TokenType `json:"refresh_token"`
}

type TokenType struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func GenerateJWT(guid string) Tokens {
	currentTime := time.Now()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &Claims{
		GUID: guid,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: currentTime.Add(1 * time.Hour).Unix(),
		},
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		GUID: guid,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: currentTime.Add(24 * time.Hour).Unix(),
		},
	})

	accessTokenString, _ := accessToken.SignedString(SigningKey)
	refreshTokenString, _ := refreshToken.SignedString(SigningKey)
	return Tokens{
		AccessToken: TokenType{
			Token:     accessTokenString,
			CreatedAt: currentTime,
			ExpiresAt: currentTime.Add(1 * time.Hour),
		},
		RefreshToken: TokenType{
			Token:     refreshTokenString,
			CreatedAt: currentTime,
			ExpiresAt: currentTime.Add(24 * time.Hour)},
	}
}

func (pairTokens Tokens) CreateDatabaseDocument(guid string) database.UserTokens {
	refreshToken := []byte(pairTokens.RefreshToken.Token)
	hashedRefreshToken, _ := bcrypt.GenerateFromPassword(refreshToken, bcrypt.DefaultCost)
	return database.UserTokens{
		GUID: guid,
		RefreshTokens: []database.TokenType{{
			Token:     string(hashedRefreshToken),
			CreatedAt: pairTokens.RefreshToken.CreatedAt,
			ExpiredAt: pairTokens.RefreshToken.ExpiresAt,
		}},
	}
}
