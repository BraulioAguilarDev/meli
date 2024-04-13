package service

import (
	"context"
	"meli/internal/core/domain"
	"meli/internal/core/port"
)

type itemService struct {
	repository port.ItemResopitory
}

func ProvideItemService(repo port.ItemResopitory) *itemService {
	return &itemService{
		repo,
	}
}

func (srv *itemService) CreateItem(ctx context.Context, item *domain.Item) (*domain.Item, error) {
	return srv.repository.CreateItem(ctx, item)
}
