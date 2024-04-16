package service

import (
	"context"
	"fmt"
	c "meli/internal/adapter/config"
	"meli/internal/adapter/reader"
	"meli/internal/core/domain"
	"meli/internal/core/port"
	"meli/pkg/melihttp"
	"sync"
)

var (
	MaxConcurrency = 10
)

type baseService struct {
	repository port.ItemResopitory
}

func ProvideBaseService(repo port.ItemResopitory) *baseService {
	return &baseService{
		repo,
	}
}

func (srv *baseService) CreateItem(ctx context.Context, item *domain.Item) (*domain.Item, error) {
	return srv.repository.CreateItem(ctx, item)
}

func (srv *baseService) UploadFile(ctx context.Context, uploadFile *domain.UploadFile) error {
	records, err := reader.ReadFileByType(uploadFile.File)
	if err != nil {
		return err
	}

	rows := make([]domain.Row, 0)
	for i, record := range records {
		if i == 0 {
			continue
		}

		rows = append(rows, domain.Row{
			Site: record[0],
			ID:   record[1],
		})
	}

	if err := srv.Queries(ctx, rows); err != nil {
		return err
	}

	return nil
}

func (srv *baseService) Queries(ctx context.Context, rows []domain.Row) error {
	fetchers := []port.QueryFetcher{
		CategoryFetcher{fmt.Sprintf("%s/categories", c.Config.API.URL)},
		CurrencyFetcher{fmt.Sprintf("%s/currencies", c.Config.API.URL)},
		SellerFetcher{fmt.Sprintf("%s/users", c.Config.API.URL)},
	}

	itmCh := make(chan domain.Item, MaxConcurrency)
	errCh := make(chan error)
	client := melihttp.NewClient()

	go fetchURL(rows, fetchers, itmCh, errCh, client)

	for {
		select {
		case result, ok := <-itmCh:
			if !ok {
				return nil
			}
			_, err := srv.repository.CreateItem(ctx, &result)
			if err != nil {
				fmt.Printf("creatting item in db error: %v\n", err.Error())
			}
		case err, ok := <-errCh:
			if !ok {
				return nil
			}
			fmt.Println("Error:", err)
		}
	}
}

func fetchURL(rows []domain.Row, fetchers []port.QueryFetcher, itmCh chan domain.Item, errCh chan error, client *melihttp.Request) {
	defer close(itmCh)

	sem := make(chan struct{}, MaxConcurrency)
	var wg sync.WaitGroup

	for _, row := range rows {
		wg.Add(1)
		go func(row domain.Row) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() {
				<-sem
			}()

			// Getting main data
			req := &ItemFetcher{
				fmt.Sprintf("%s/items", c.Config.API.URL),
			}

			fmt.Printf("Getting item from %s with ID: %v\n", c.Config.API.URL, row.ID)
			response, err := req.Fetch(client, map[string]string{
				"site": row.Site, "id": row.ID,
			})
			if err != nil {
				errCh <- err
			}

			result := domain.Item{
				ID:        row.ID,
				Site:      row.Site,
				StartTime: response["date_created"],
				Price:     response["price"],
			}

			for _, f := range fetchers {
				res, err := f.Fetch(client, response)
				if err != nil {
					errCh <- err
				}

				filling(res, &result)
			}

			itmCh <- result
		}(row)
	}

	wg.Wait()
}

func filling(fetch map[string]string, result *domain.Item) *domain.Item {
	if name, ok := fetch["name"]; ok {
		result.Name = name // category
	}

	if desc, ok := fetch["description"]; ok {
		result.Description = desc // currency
	}

	if nick, ok := fetch["nickname"]; ok {
		result.Nickname = nick // seller
	}

	return result
}
