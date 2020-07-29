package database

type UserTokens struct {
	GUID          string `bson:"_id,omitempty"`
	IP            string `bson:"ip,omitempty"`
	OS            string `bson:"os,omitempty"`
	UserAgent     string `bson:"user_agent"`
	RefreshTokens string `bson:"refresh_tokens,omitempty"`
	CreatedAt     string `bson:"created_at, omitempty"`
	ExpiredAt     string `bson:"expired_at"`
}
