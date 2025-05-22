package models

// Основная модель цитаты
type Quote struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Text   string `json:"quote"`
}

// Для запроса на создание циаты
type CreateQuoteRequest struct {
	Author string `json:"author"`
	Text   string `json:"quote"`
}
