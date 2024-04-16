package port

import (
	"meli/pkg/melihttp"
)

// QueryFetcher is an interface for interacting calls to external API
type QueryFetcher interface {
	Fetch(client *melihttp.Request, row map[string]string) (map[string]string, error)
}
