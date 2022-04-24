package repository

import (
	"context"
	"transaction-service/models"
)

type DBOps interface {
	CreateAccount(ctx context.Context, profile *models.Account) error
	GetAccount(ctx context.Context, id int64) (*models.Account, error)
	GetAccountByDocument(ctx context.Context, docNo string) (*models.Account, error)
	GetOperationType(ctx context.Context, id int64) (*models.Operations, error)
	CreateTransaction(ctx context.Context, txn *models.Transaction) error
}
