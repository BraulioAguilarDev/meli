package service

import (
	"encoding/json"
	"fmt"
	"io"
	"meli/pkg/melihttp"
	"net/http"
)

type CategoryFetcher struct {
	Path string
}

type CategoryResponse struct {
	Name string `json:"name"`
}

func (caf CategoryFetcher) Fetch(client *melihttp.Request, row map[string]string) (map[string]string, error) {
	options := &melihttp.Options{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("%s/%s", caf.Path, row["category_id"]),
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

	var response CategoryResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return map[string]string{
		"name": response.Name,
	}, nil
}
