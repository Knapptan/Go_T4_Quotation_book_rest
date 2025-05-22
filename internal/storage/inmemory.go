package storage

import (
	"context"
	"sync"

	"github.com/Knapptan/Go_T1_Name_info_rest/internal/models"
	"github.com/google/uuid"
)

type InMemoryStorage struct {
	quotes map[string]models.Quote
	mu     sync.RWMutex
}

func (s *InMemoryStorage) Create(ctx context.Context, quoteReq models.CreateQuoteRequest) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New().String()
	quote := models.Quote{
		ID:     id,
		Author: quoteReq.Author,
		Text:   quoteReq.Text,
	}
	s.quotes[id] = quote
	return id, nil
}
