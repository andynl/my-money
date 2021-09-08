package entity

import "time"

type Transaction struct {
	ID         string
	Type       string
	CategoryID string
	Name       string
	Amount     float64
	CreatedAt  time.Time
}
