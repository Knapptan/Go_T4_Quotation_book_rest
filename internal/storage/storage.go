package storage

import "context"

type QuotesRepository interface {
	CreateQuote(ctx context.Context) error
	GetQuote(ctx context.Context) error
	GetAllQuote(ctx context.Context) error
	UpdateQuote(ctx context.Context) error
	DeleteQuote(ctx context.Context) error
}
