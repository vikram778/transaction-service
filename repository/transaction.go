package repository

import (
	"context"
	"go.uber.org/zap"
	"time"
	"transaction-service/models"
	"transaction-service/pkg/log"
)

func (r *dBOps) CreateTransaction(ctx context.Context, txn *models.Transaction) error {

	txn.EventDate = time.Now()
	_, err := r.db.Exec(
		createTransactionQuery,
		txn.AccountID,
		txn.OperationTypeID,
		txn.Amount,
		txn.EventDate,
	)

	if err != nil {
		log.Info("Create transaction error", zap.Error(err))
		return err
	}

	return nil
}
