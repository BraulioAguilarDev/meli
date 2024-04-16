package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	c "meli/internal/adapter/config"
	"meli/pkg/melihttp"
	"net/http"
	"time"
)

// Implementing for /items API
type ItemFetcher struct {
	Path string
}

type ItemResponse struct {
	Code int `json:"code"`
	Body struct {
		ID          string    `json:"id,omitempty"`
		Price       float64   `json:"price,omitempty"`
		DateCreated time.Time `json:"date_created,omitempty"` // smart_date key doesn't exists
		SellerID    int       `json:"seller_id,omitempty"`
		CategoryID  string    `json:"category_id,omitempty"`
		CurrencyID  string    `json:"currency_id,omitempty"`
	}
}

// Fetch executes call to /items then fills at ItemResponse
func (itf ItemFetcher) Fetch(client *melihttp.Request, row map[string]string) (map[string]string, error) {
	clave := fmt.Sprintf("%s%s", row["site"], row["id"])

	options := &melihttp.Options{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("%s?ids=%s", itf.Path, clave),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", c.Config.API.TOKEN),
		},
	}

	res, err := client.MakeRequest(options)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result []ItemResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("not results")
	}

	return map[string]string{
		"price":        fmt.Sprintf("%v", result[0].Body.Price),
		"date_created": result[0].Body.DateCreated.String(),
		"category_id":  result[0].Body.CategoryID,
		"currency_id":  result[0].Body.CurrencyID,
		"seller_id":    fmt.Sprintf("%d", result[0].Body.SellerID),
	}, nil
}
