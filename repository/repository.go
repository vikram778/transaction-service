package repository

import (
	"context"
	"transaction-service/models"
)

type DBOps interface {
	CreateAccount(ctx context.Context, profile *models.Account) (*models.Account, error)
	GetAccount(ctx context.Context, id int64) (*models.Account, error)
	CreateTransaction(ctx context.Context, txn *models.Transaction) error
}
