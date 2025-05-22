package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Knapptan/Go_T1_Name_info_rest/internal/models"
	"github.com/Knapptan/Go_T1_Name_info_rest/internal/storage"
	"github.com/gorilla/mux"
)

type QuoteHandler struct {
	repo storage.QuotesRepository
}

func NewQuoteHandler(repo storage.QuotesRepository) *QuoteHandler {
	return &QuoteHandler{repo: repo}
}

func (h *QuoteHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var request models.CreateQuoteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	quote := models.CreateQuoteRequest{
		Author: request.Author,
		Text:   request.Text,
	}

	id, err := h.repo.Create(r.Context(), quote)
	if err != nil {
		http.Error(w, "Failed to create quote", http.StatusInternalServerError)
		return
	}

	log.Printf("CreateQuote: id %s", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *QuoteHandler) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	quotes, err := h.repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to get all quotes", http.StatusInternalServerError)
		return
	}

	log.Printf("GetAllQuotes: len quotes %d", len(quotes))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(quotes); err != nil {
		http.Error(w, "Failed to encode quotes", http.StatusInternalServerError)
	}
}

func (h *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.repo.GetRandom(r.Context())
	if err != nil {
		http.Error(w, "Failed to get randome quote", http.StatusInternalServerError)
		return
	}

	log.Printf("GetRandomQuote: id %s", quote.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(quote); err != nil {
		http.Error(w, "Failed to encode quote", http.StatusInternalServerError)
	}
}

func (h *QuoteHandler) GetByAuthorQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	if author == "" {
		http.Error(w, "'author' query parameter is required", http.StatusBadRequest)
		return
	}

	quotes, err := h.repo.GetByAuthor(r.Context(), author)
	if err != nil {
		http.Error(w, "Failed to fetch quotes for author", http.StatusInternalServerError)
		return
	}

	log.Printf("GetByAuthorQuotes: author %s", author)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(quotes); err != nil {
		http.Error(w, "Failed to encode quotes", http.StatusInternalServerError)
	}
}

func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.repo.Delete(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrNotFound):
			http.Error(w, "Quote not found", http.StatusNotFound)
		default:
			log.Printf("Delete error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	log.Printf("DeleteQuote: id %s", id)

	w.WriteHeader(http.StatusNoContent)
}
