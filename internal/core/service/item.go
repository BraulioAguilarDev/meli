package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"meli/pkg/melihttp"
	"net/http"
	"time"
)

type ItemFetcher struct {
	Path string
}

type ItemResponse struct {
	Code int `json:"code"`
	Body struct {
		ID          string    `json:"id"`
		Price       float64   `json:"price"`
		DateCreated time.Time `json:"date_created"` // smart_date key doesn't exists
		SellerID    int       `json:"seller_id"`
		CategoryID  string    `json:"category_id"`
		CurrencyID  string    `json:"currency_id"`
	}
}

func (itf ItemFetcher) Fetch(client *melihttp.Request, row map[string]string) (map[string]string, error) {
	clave := fmt.Sprintf("%s%s", row["site"], row["id"])

	options := &melihttp.Options{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("%s?ids=%s", itf.Path, clave),
		Headers: map[string]string{
			"Authorization": "Bearer APP_USR-7032346726927327-041520-c3f86dcd27a37a83df54a11a1ecf28b2-654966372",
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
