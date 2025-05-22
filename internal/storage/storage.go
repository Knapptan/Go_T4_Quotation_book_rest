package storage

import (
	"context"

	"github.com/Knapptan/Go_T1_Name_info_rest/internal/models"
)

type QuotesRepository interface {
	Create(ctx context.Context, quote models.CreateQuoteRequest) (string, error)
	GetAll(ctx context.Context) ([]models.Quote, error)
	GetRandom(ctx context.Context) (models.Quote, error)
	GetByAuthor(ctx context.Context, author string) ([]models.Quote, error)
	Delete(ctx context.Context, id string) error
}
