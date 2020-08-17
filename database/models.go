package database

import (
	"time"
)

type UserTokens struct {
	GUID          string      `json:"guid" bson:"_id,omitempty"`
	RefreshTokens []TokenType `json:"refresh_tokens" bson:"refresh_tokens"`
}

type TokenType struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	ExpiredAt time.Time `json:"expired_at" bson:"expired_at"`
}
