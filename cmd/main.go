package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Knapptan/Go_T1_Name_info_rest/internal/config"
	"github.com/Knapptan/Go_T1_Name_info_rest/internal/handler"
	"github.com/Knapptan/Go_T1_Name_info_rest/internal/storage"
	"github.com/gorilla/mux"
)

func main() {

	repo := storage.NewInMemoryStorage()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Config load error")
	}

	quoteHandler := handler.NewQuoteHandler(repo)

	r := mux.NewRouter()
	r.HandleFunc("/quotes", quoteHandler.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes", quoteHandler.GetAllQuotes).Methods("GET")
	r.HandleFunc("/quotes/random", quoteHandler.GetRandomQuote).Methods("GET")
	r.HandleFunc("/quotes", quoteHandler.GetByAuthorQuotes).Methods("GET")
	r.HandleFunc("/quotes/{id}", quoteHandler.DeleteQuote).Methods("DELETE")

	address := cfg.StrAdress()

	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server started on adress: %s", address)
	serverErr := srv.ListenAndServe()
	if serverErr != nil && serverErr != http.ErrServerClosed {
		log.Fatal(serverErr)
	}
}
