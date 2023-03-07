package models

import (
	"time"
)

type NSELedgersLogs struct {
	Id            int64     `json:"id"`
	ClientCode    string    `json:"client_code"`
	Segment       string    `json:"segment"`
	Selection     string    `json:"selection"`
	Amount        float64   `json:"amount"`
	PendingAmount float64   `json:"pending_amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
