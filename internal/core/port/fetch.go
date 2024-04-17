package port

import (
	"meli/pkg/melihttp"
)

//go:generate mockgen -source=fetch.go -destination=mock/fetch.go -package=mock

// QueryFetcher is an interface for interacting calls to external API
type QueryFetcher interface {
	Fetch(client *melihttp.Request, row map[string]string) (map[string]string, error)
}
