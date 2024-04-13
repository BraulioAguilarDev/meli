package port

import (
	"context"
	"meli/internal/core/domain"
)

type ItemResopitory interface {
	CreateItem(context.Context, *domain.Item) (*domain.Item, error)
	// CreateItems(context.Context, []domain.Item) error
}

type ItemService interface {
	CreateItem(context.Context, *domain.Item) (*domain.Item, error)
}
