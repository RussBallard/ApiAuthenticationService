package handler

// Handler - агрегатор обработчиков
type Handler interface {
	CreatePairTokens()
	UpdatePairTokens()
	DeleteOneToken()
	DeleteAllTokens()
}
