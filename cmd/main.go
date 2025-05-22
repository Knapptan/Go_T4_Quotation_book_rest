package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Knapptan/Go_T1_Name_info_rest/internal/handler"
	"github.com/Knapptan/Go_T1_Name_info_rest/internal/storage"
	"github.com/gorilla/mux"
)

// Добавление новой цитаты (POST /quotes)
// Получение всех цитат (GET /quotes)
// Получение случайной цитаты (GET /quotes/random)
// Фильтрация по автору (GET /quotes?author=Confucius)
// Удаление цитаты по ID (DELETE /quotes/{id})

func main() {

	repo := storage.NewInMemoryStorage()

	quoteHandler := handler.NewQuoteHandler(repo)

	r := mux.NewRouter()
	r.HandleFunc("/quotes", quoteHandler.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes", quoteHandler.GetAllQuotes).Methods("GET")
	r.HandleFunc("/quotes/random", quoteHandler.GetRandomQuote).Methods("GET")
	r.HandleFunc("/quotes", quoteHandler.GetByAuthorQuotes).Methods("GET")
	r.HandleFunc("/quotes/{id}", quoteHandler.DeleteQuote).Methods("DELETE")
	

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
