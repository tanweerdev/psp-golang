package models

import (
	"time"
)

type Transaction struct {
	ID        string    `json:"id"`
	Payment   Payment   `json:"payment"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
