package handler

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// BaseHandler - обработчик
type BaseHandler struct {
	MongoDB *mongo.Database
}

// NewBaseHandler - конструктор обработчиков
func NewBaseHandler() (_ *BaseHandler, err error) {
	handler := new(BaseHandler)

	return handler, nil
}

// CreatePairTokens - uri для создания access и refresh токена
func (b *BaseHandler) CreatePairTokens(w http.ResponseWriter, r *http.Request) {

}

// UpdatePairTokens - uri для обновления access и refresh токена
func (b *BaseHandler) UpdatePairTokens(w http.ResponseWriter, r *http.Request) {

}

// DeleteOneToken - uri для удаления одного конкретного токена
func (b *BaseHandler) DeleteOneToken(w http.ResponseWriter, r *http.Request) {

}

// DeleteAllTokens - uri для удаления всех токенов
func (b *BaseHandler) DeleteAllTokens(w http.ResponseWriter, r *http.Request) {

}
