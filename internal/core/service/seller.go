package service

import (
	"encoding/json"
	"fmt"
	"io"
	c "meli/internal/adapter/config"
	"meli/pkg/melihttp"
	"net/http"
)

// Implementing for /users API
type SellerFetcher struct {
	Path string
}

type SellerResponse struct {
	Nickname string `json:"nickname,omitempty"`
}

// Fetch executes call to /users then fills at SellerResponse
func (sef SellerFetcher) Fetch(client *melihttp.Request, row map[string]string) (map[string]string, error) {
	options := &melihttp.Options{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("%s/%s", sef.Path, row["seller_id"]),
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

	var response SellerResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return map[string]string{
		"nickname": response.Nickname,
	}, nil
}
