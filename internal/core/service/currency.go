package service

import (
	"encoding/json"
	"fmt"
	"io"
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

	var currency CurrencyResponse
	if err = json.Unmarshal(body, &currency); err != nil {
		return nil, err
	}

	return map[string]string{
		"description": currency.Description,
	}, nil
}
