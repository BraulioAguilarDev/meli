package port

import (
	"context"
	"meli/internal/core/domain"
)

//go:generate mockgen -source=item.go -destination=mock/item.go -package=mock

// ItemResopitory is an interface for interacting with item-related data
type ItemResopitory interface {
	CreateItem(context.Context, *domain.Item) (*domain.Item, error)
}

// ItemService is an interface for interacting with item-related business logic
type ItemService interface {
	CreateItem(context.Context, *domain.Item) (*domain.Item, error)
	UploadFile(context.Context, *domain.UploadFile) error
}
