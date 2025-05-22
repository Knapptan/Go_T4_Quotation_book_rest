package storage

import (
	"context"
	"fmt"
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

func (s *InMemoryStorage) GetAll(ctx context.Context) ([]models.Quote, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	quotes := make([]models.Quote, 0, len(s.quotes))
	for _, q := range s.quotes {
		quotes = append(quotes, q)
	}
	return quotes, nil
}

func (s *InMemoryStorage) GetRandom(ctx context.Context) (models.Quote, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var quote models.Quote
	for _, q := range s.quotes {
		quote = q
		break
	}
	return quote, nil
}

func (s *InMemoryStorage) GetByAuthor(ctx context.Context, author string) ([]models.Quote, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	quotes := make([]models.Quote, 0, 1)
	for _, q := range s.quotes {
		if q.Author == author {
			quotes = append(quotes, q)
		}
	}
	return quotes, nil
}

func (s *InMemoryStorage) Delete(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.quotes[id]; !exists {
		return fmt.Errorf("quote with id %q not found", id)
	}

	delete(s.quotes, id)
	return nil
}
