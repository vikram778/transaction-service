package models

type Operations struct {
	OperationTypeID string `json:"operation_type_id" db:"operation_type_id"`
	Description     string `json:"description" db:"description"`
}
