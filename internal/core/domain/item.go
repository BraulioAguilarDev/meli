package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Item struct {
	ID          uint
	Site        string
	Price       decimal.Decimal
	StartTime   time.Time
	Name        string
	Description string
	Nickname    string
}
