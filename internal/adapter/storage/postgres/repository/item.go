package repository

import (
	"context"
	"meli/internal/core/domain"
	"strconv"
)

/**
 * itemRepository implements port.ItemResopitory interface
 * and provides an access to the postgres database by sqlc
 */
type itemRepository struct {
	queries *Queries
}

// NewItemRepository creates a new item repository instance
func NewItemRepository(queries *Queries) *itemRepository {
	return &itemRepository{
		queries,
	}
}

// CreateItem creates a new item record in the database
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
