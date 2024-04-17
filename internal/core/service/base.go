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
	MaxConcurrency = 10 // Maximum number of concurrent goroutines
)

/**
 * baseService implements port.ItemService interface
 * and provides an access to the item repository
 */
type baseService struct {
	repository port.ItemResopitory
}

// ProvideBaseService creates a new item service instance
func ProvideBaseService(repo port.ItemResopitory) *baseService {
	return &baseService{
		repo,
	}
}

// CreateItem saves an item
func (srv *baseService) CreateItem(ctx context.Context, item *domain.Item) (*domain.Item, error) {
	return srv.repository.CreateItem(ctx, item)
}

// UploadFile process a file then save it if all is ok
func (srv *baseService) UploadFile(ctx context.Context, uploadFile *domain.UploadFile) error {

	// ReadFileByType executes the reading according file received
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

	/*
	 Pass `rows` data to Queries function for calling APIs
	 Item, Category, Seller, Currency

	 Item is the main request becouse then we need info as category_id, currency_id, etc
	*/
	if err := srv.Queries(ctx, rows); err != nil {
		return err
	}

	return nil
}

// Queries call external API by goroutines
func (srv *baseService) Queries(ctx context.Context, rows []domain.Row) error {
	fetchers := []port.QueryFetcher{
		CategoryFetcher{fmt.Sprintf("%s/categories", c.Config.API.URL)},
		CurrencyFetcher{fmt.Sprintf("%s/currencies", c.Config.API.URL)},
		SellerFetcher{fmt.Sprintf("%s/users", c.Config.API.URL)},
		// Adding new queryFetch...
	}

	itmCh := make(chan domain.Item, MaxConcurrency)
	errCh := make(chan error)
	client := melihttp.NewClient()

	go fetchURL(rows, fetchers, itmCh, errCh, client)

	// Reading channels
	for {
		select {
		case result, ok := <-itmCh:
			if !ok {
				return nil
			}
			// Saving a item in db
			_, err := srv.repository.CreateItem(ctx, &result) // TODO: Save in bulk
			if err != nil {
				fmt.Printf("creatting item in db error: %v\n", err.Error())
			}
		case err, ok := <-errCh:
			if !ok {
				return nil
			}
			fmt.Println("Error:", err) // TODO: Adding retry
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

			sem <- struct{}{} // Acquire a semaphore
			defer func() {
				<-sem // release semaphore
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
