package storage_test

import (
	"context"
	"strconv"
	"sync"
	"testing"

	"github.com/Knapptan/Go_T1_Name_info_rest/internal/models"
	"github.com/Knapptan/Go_T1_Name_info_rest/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestInMemoryStorage_Create(t *testing.T) {
	t.Run("successful creation", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		quote := models.CreateQuoteRequest{Author: "Test", Text: "Test quote"}

		id, err := s.Create(context.Background(), quote)
		require.NoError(t, err)
		require.NotEmpty(t, id)
	})

	t.Run("ID uniqueness", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		quote := models.CreateQuoteRequest{Author: "Test", Text: "Test quote"}

		id1, _ := s.Create(context.Background(), quote)
		id2, _ := s.Create(context.Background(), quote)
		require.NotEqual(t, id1, id2)
	})
}

func TestInMemoryStorage_GetAll(t *testing.T) {
	t.Run("empty storage", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		quotes, err := s.GetAll(context.Background())
		require.NoError(t, err)
		require.Empty(t, quotes)
	})

	t.Run("with data", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		expected := []models.CreateQuoteRequest{
			{Author: "A1", Text: "Q1"},
			{Author: "A2", Text: "Q2"},
		}

		for _, q := range expected {
			s.Create(context.Background(), q)
		}

		actual, err := s.GetAll(context.Background())
		require.NoError(t, err)
		require.Len(t, actual, len(expected))

		// Проверяем что ID заполнены
		for _, q := range actual {
			require.NotEmpty(t, q.ID)
		}
	})
}

func TestInMemoryStorage_Delete(t *testing.T) {
	t.Run("existing quote", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		id, _ := s.Create(context.Background(), models.CreateQuoteRequest{Author: "Test", Text: "Test"})

		err := s.Delete(context.Background(), id)
		require.NoError(t, err)

		quotes, _ := s.GetAll(context.Background())
		require.Empty(t, quotes)
	})

	t.Run("non-existent quote", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		err := s.Delete(context.Background(), "invalid_id")
		require.ErrorIs(t, err, storage.ErrNotFound)
	})
}

func TestInMemoryStorage_Concurrency(t *testing.T) {
	s := storage.NewInMemoryStorage()
	var wg sync.WaitGroup

	// Конкурентные записи
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := s.Create(context.Background(), models.CreateQuoteRequest{Author: "Concurrent", Text: "Test"})
			require.NoError(t, err)
		}()
	}

	// Конкурентные чтения
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := s.GetAll(context.Background())
			require.NoError(t, err)
		}()
	}

	wg.Wait()
	quotes, _ := s.GetAll(context.Background())
	require.Len(t, quotes, 100)
}

func TestInMemoryStorage_GetByAuthor(t *testing.T) {
	s := storage.NewInMemoryStorage()
	testData := []models.CreateQuoteRequest{
		{Author: "A1", Text: "Q1"},
		{Author: "A1", Text: "Q2"},
		{Author: "A2", Text: "Q3"},
	}

	for _, q := range testData {
		s.Create(context.Background(), q)
	}

	t.Run("existing author", func(t *testing.T) {
		result, err := s.GetByAuthor(context.Background(), "A1")
		require.NoError(t, err)
		require.Len(t, result, 2)
	})

	t.Run("non-existent author", func(t *testing.T) {
		result, err := s.GetByAuthor(context.Background(), "Unknown")
		require.NoError(t, err)
		require.Empty(t, result)
	})
}

func TestInMemoryStorage_GetRandom(t *testing.T) {
	t.Run("empty storage", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		_, err := s.GetRandom(context.Background())
		require.ErrorIs(t, err, storage.ErrNotFound)
	})

	t.Run("with single quote", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		test := models.CreateQuoteRequest{Author: "Test", Text: "Test"}
		id, _ := s.Create(context.Background(), test)
		idS, _ := strconv.Atoi(id)

		expected := models.Quote{
			ID:     idS,
			Author: test.Author,
			Text:   test.Text,
		}

		actual, err := s.GetRandom(context.Background())
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("with multiple quotes", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		quotes := []models.CreateQuoteRequest{
			{Author: "A1", Text: "Q1"},
			{Author: "A2", Text: "Q2"},
		}

		for _, q := range quotes {
			s.Create(context.Background(), q)
		}

		actual, err := s.GetRandom(context.Background())
		require.NoError(t, err)
		require.Contains(t, quotes, actual)
	})
}
