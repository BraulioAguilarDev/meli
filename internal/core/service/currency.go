package service

import (
	"encoding/json"
	"fmt"
	"io"
	c "meli/internal/adapter/config"
	"meli/pkg/melihttp"
	"net/http"
)

type CurrencyFetcher struct {
	Path string
}

type CurrencyResponse struct {
	ID            string `json:"id,omitempty"`
	Symbol        string `json:"symbol,omitempty"`
	Description   string `json:"description,omitempty"`
	DecimalPlaces int    `json:"decimal_places,omitempty"`
}

func (cuf CurrencyFetcher) Fetch(client *melihttp.Request, row map[string]string) (map[string]string, error) {
	options := &melihttp.Options{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("%s/%s", cuf.Path, row["currency_id"]),
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

	var currency CurrencyResponse
	if err = json.Unmarshal(body, &currency); err != nil {
		return nil, err
	}

	return map[string]string{
		"description": currency.Description,
	}, nil
}
