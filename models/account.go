package models

import "time"

type Account struct {
	AccountID      int64     `json:"account_id" db:"account_id"`
	DocumentNumber string    `json:"document_number" db:"document_number"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
