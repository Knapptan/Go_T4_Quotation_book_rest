package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Knapptan/Go_T1_Name_info_rest/internal/handler"
	"github.com/Knapptan/Go_T1_Name_info_rest/internal/models"
	"github.com/Knapptan/Go_T1_Name_info_rest/internal/storage"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

type MockStorage struct {
	CreateFunc      func(context.Context, models.CreateQuoteRequest) (string, error)
	GetAllFunc      func(context.Context) ([]models.Quote, error)
	GetRandomFunc   func(ctx context.Context) (models.Quote, error)
	GetByAuthorFunc func(ctx context.Context, author string) ([]models.Quote, error)
	DeleteFunc      func(context.Context, string) error
}

func (m *MockStorage) Create(ctx context.Context, q models.CreateQuoteRequest) (string, error) {
	return m.CreateFunc(ctx, q)
}

func (m *MockStorage) GetAll(ctx context.Context) ([]models.Quote, error) {
	return m.GetAllFunc(ctx)
}

func (m *MockStorage) GetRandom(ctx context.Context) (models.Quote, error) {
	return m.GetRandomFunc(ctx)
}

func (m *MockStorage) GetByAuthor(ctx context.Context, author string) ([]models.Quote, error) {
	return m.GetByAuthorFunc(ctx, author)
}

func (m *MockStorage) Delete(ctx context.Context, id string) error {
	return m.DeleteFunc(ctx, id)
}

func TestCreateQuoteHandler(t *testing.T) {
	mockStorage := &MockStorage{
		CreateFunc: func(ctx context.Context, req models.CreateQuoteRequest) (string, error) {
			require.Equal(t, "Author", req.Author)
			require.Equal(t, "Text", req.Text)
			return "42", nil
		},
	}

	h := handler.NewQuoteHandler(mockStorage)

	body := `{"author":"Author","text":"Text"}`
	req := httptest.NewRequest(http.MethodPost, "/quotes", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	router := mux.NewRouter()
	router.HandleFunc("/quotes", h.CreateQuote).Methods(http.MethodPost)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code, "ожидаем 201 Created")

	var resp map[string]string
	err := json.NewDecoder(bytes.NewBuffer(rr.Body.Bytes())).Decode(&resp)
	require.NoError(t, err, "должен быть корректный JSON")
	require.Equal(t, "42", resp["id"], "ID в ответе должен совпадать с тем, что вернул репозиторий")
}

func TestDeleteQuote_NotFound(t *testing.T) {
	mockStorage := &MockStorage{
		DeleteFunc: func(ctx context.Context, id string) error {
			require.Equal(t, "123", id)
			return storage.ErrNotFound
		},
	}
	h := handler.NewQuoteHandler(mockStorage)

	router := mux.NewRouter()
	router.HandleFunc("/quotes/{id}", h.DeleteQuote).Methods(http.MethodDelete)

	req := httptest.NewRequest(http.MethodDelete, "/quotes/123", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusNotFound, rr.Code)
}
