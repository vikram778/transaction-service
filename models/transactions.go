package models

import "time"

type Transaction struct {
	TransactionID   string    `json:"transaction_id" db:"transaction_id"`
	AccountID       int64     `json:"account_id" db:"account_id"`
	OperationTypeID int64     `json:"operation_type_id" db:"operation_type_id"`
	Amount          float64   `json:"amount" db:"amount"`
	EventDate       time.Time `json:"event_date" db:"event_date"`
}
