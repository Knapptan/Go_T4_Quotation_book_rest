package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Knapptan/Go_T1_Name_info_rest/internal/models"
	"github.com/Knapptan/Go_T1_Name_info_rest/internal/storage"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
