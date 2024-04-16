package repository

import (
	"context"
	"meli/internal/core/domain"
	"strconv"
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
	id, _ := strconv.Atoi(item.ID)
	_, err := repo.queries.CreateItem(ctx, CreateItemParams{
		ID:          id,
		Site:        item.Site,
		Price:       item.Price,
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
