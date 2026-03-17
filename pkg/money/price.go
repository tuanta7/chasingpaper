package money

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Currency string

const (
	USD Currency = "usd"
	EUR Currency = "eur"
	SGD Currency = "sgd"
	VND Currency = "vnd"
)

type Price struct {
	Amount   float64  `json:"amount"`
	Currency Currency `json:"currency"`
}

// Scan implements the Scanner interface.
func (p *Price) Scan(src any) error {
	if src == nil {
		return nil
	}

	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, p)
	case string:
		return json.Unmarshal([]byte(v), p)
	default:
		return errors.New("type assertion to []byte or string failed")
	}
}

// Value implements the driver Valuer interface.
func (p Price) Value() (driver.Value, error) {
	return json.Marshal(p)
}
