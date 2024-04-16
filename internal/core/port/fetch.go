package port

import (
	"meli/pkg/melihttp"
)

type QueryFetcher interface {
	Fetch(client *melihttp.Request, id map[string]string) (map[string]string, error)
}
