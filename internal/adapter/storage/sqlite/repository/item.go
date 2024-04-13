package repository

import (
	"context"
	"meli/internal/core/domain"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
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

	itemCreated, err := repo.queries.CreateItem(ctx, CreateItemParams{
		ID:          int32(item.ID),
		Site:        item.Site,
		Price:       item.Price.String(),
		SmartTime:   pgtype.Timestamp{Time: time.Now()},
		Name:        item.Name,
		Description: item.Description,
		Nickname:    item.Nickname,
	})

	if err != nil {
		return nil, err
	}

	item.StartTime = itemCreated.SmartTime.Time
	return item, nil
}
