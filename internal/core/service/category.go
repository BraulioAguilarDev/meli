package service

import (
	"encoding/json"
	"fmt"
	"io"
	c "meli/internal/adapter/config"
	"meli/pkg/melihttp"
	"net/http"
)

type CategoryFetcher struct {
	Path string
}

type CategoryResponse struct {
	Name string `json:"name,omitempty"`
}

func (caf CategoryFetcher) Fetch(client *melihttp.Request, row map[string]string) (map[string]string, error) {
	options := &melihttp.Options{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("%s/%s", caf.Path, row["category_id"]),
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

	var response CategoryResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return map[string]string{
		"name": response.Name,
	}, nil
}
