package repository

import (
	"context"
	"meli/internal/core/domain"
)

type itemRepository struct {
	queries *Queries
}

func NewItemRepository(queries *Queries) *itemRepository {
	return &itemRepository{
		queries,
	}
}

func (repo *itemRepository) CreateItem(ctx context.Context, item *domain.Item) (*domain.Item, error) {

	_, err := repo.queries.CreateItem(ctx, CreateItemParams{
		ID:          item.ID,
		Site:        item.Site,
		Price:       item.Price.String(),
		SmartTime:   item.StartTime,
		Name:        item.Name,
		Description: item.Description,
		Nickname:    item.Nickname,
	})

	if err != nil {
		return nil, err
	}

	return item, nil
}
