package service

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"meli/internal/core/domain"
	"meli/internal/core/port"
	"meli/pkg/melihttp"
	"sync"

	"github.com/shopspring/decimal"
)

var (
	MaxConcurrency = 10
)

type baseService struct {
	repository port.ItemResopitory
	url        string
}

func ProvideBaseService(repo port.ItemResopitory, url string) *baseService {
	return &baseService{
		repo, url,
	}
}

func (srv *baseService) CreateItem(ctx context.Context, item *domain.Item) (*domain.Item, error) {
	return srv.repository.CreateItem(ctx, item)
}

func (srv *baseService) UploadFile(ctx context.Context, uploadFile *domain.UploadFile) error {
	// TODO: multiple formats support
	if uploadFile.File.Header.Get("Content-Type") != "text/csv" {
		return errors.New("file is not a CSV file")
	}

	file, err := uploadFile.File.Open()
	if err != nil {
		return err
	}

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return err
	}

	rows := make([]domain.Row, 0)
	for _, row := range data {
		rows = append(rows, domain.Row{
			Site: row[0],
			ID:   row[1],
		})
	}

	if err := srv.Queries(ctx, rows); err != nil {
		return err
	}

	return nil
}

func (srv *baseService) Queries(ctx context.Context, rows []domain.Row) error {
	// Call external services
	fetchers := []port.QueryFetcher{
		CategoryFetcher{fmt.Sprintf("%s/categories", srv.url)},
		CurrencyFetcher{fmt.Sprintf("%s/currencies", srv.url)},
		SellerFetcher{fmt.Sprintf("%s/users", srv.url)},
		// Adds other services...
	}

	ch := make(chan domain.Item, MaxConcurrency)
	client := melihttp.NewClient()

	go fetchURL(rows, fetchers, ch, client, srv.url)

	for item := range ch {
		fmt.Println("Item for saving:", item)
	}

	return nil
}

func fetchURL(rows []domain.Row, fetchers []port.QueryFetcher, ch chan domain.Item, client *melihttp.Request, apiURL string) {
	defer close(ch)

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

			// Getting item data
			req := &ItemFetcher{
				fmt.Sprintf("%s/items", apiURL),
			}

			response, err := req.Fetch(client, map[string]string{
				"site": row.Site, "id": row.ID,
			})
			if err != nil {
				log.Printf("fetching item error: %v", err)
			}

			price, _ := decimal.NewFromString(response["price"])
			result := domain.Item{
				ID:        row.ID,
				Site:      row.Site,
				StartTime: response["date_created"],
				Price:     price,
			}

			for _, f := range fetchers {
				res, err := f.Fetch(client, response)
				if err != nil {
					log.Printf("fetching error: %v\n", err)
				}

				filling(res, &result)
			}

			ch <- result
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
