package main

import (
	"github.com/Knapptan/Go_T1_Name_info_rest/internal/handler"
	"github.com/Knapptan/Go_T1_Name_info_rest/internal/models"
	"github.com/Knapptan/Go_T1_Name_info_rest/internal/storage"
)

// Добавление новой цитаты (POST /quotes)
// Получение всех цитат (GET /quotes)
// Получение случайной цитаты (GET /quotes/random)
// Фильтрация по автору (GET /quotes?author=Confucius)
// Удаление цитаты по ID (DELETE /quotes/{id})

func main() {

	repo := storage.NewInMemoryStorage()

	quoteHandler := handler.NewQuoteHandler(repo)

	
}
