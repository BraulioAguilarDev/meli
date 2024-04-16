package domain

import (
	"github.com/shopspring/decimal"
)

type Item struct {
	ID          string
	Site        string
	Price       decimal.Decimal
	StartTime   string
	Name        string
	Description string
	Nickname    string
}
