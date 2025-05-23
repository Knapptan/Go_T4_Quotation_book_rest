package storage

import (
	"context"
	"strconv"
	"sync"

	"github.com/Knapptan/Go_T1_Name_info_rest/internal/models"
)

type InMemoryStorage struct {
	mu     sync.RWMutex
	quotes map[int]models.Quote
	nextID int
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		quotes: make(map[int]models.Quote),
		nextID: 1,
	}
}

func (s *InMemoryStorage) Create(ctx context.Context, req models.CreateQuoteRequest) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID
	s.nextID++

	quote := models.Quote{
		ID:     id,
		Author: req.Author,
		Text:   req.Text,
	}
	s.quotes[id] = quote

	return strconv.Itoa(id), nil
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

	if len(s.quotes) == 0 {
		return models.Quote{}, ErrNotFound
	}

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

	idNum, err := strconv.Atoi(id)
	if err != nil {
		return ErrNotFound
	}

	if _, exists := s.quotes[idNum]; !exists {
		return ErrNotFound
	}

	delete(s.quotes, idNum)
	return nil
}
