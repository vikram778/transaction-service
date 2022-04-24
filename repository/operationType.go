package repository

import (
	"context"
	"go.uber.org/zap"
	"transaction-service/models"
	"transaction-service/pkg/log"
)

func (r *dBOps) GetOperationType(ctx context.Context, id int64) (*models.Operations, error) {

	var operationType models.Operations

	if err := r.db.Get(&operationType, getOperationTypeQuery, id); err != nil {
		log.Info("fetch operation type error", zap.Error(err))
		return nil, err
	}
	return &operationType, nil
}
