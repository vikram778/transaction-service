package repository

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	"time"
	"transaction-service/models"
	"transaction-service/pkg/log"
)

func (r *dBOps) CreateAccount(ctx context.Context, profile *models.Account) error {

	profile.CreatedAt = time.Now()

	_, err := r.db.Exec(
		createAccountQuery,
		profile.DocumentNumber,
		profile.CreatedAt,
	)

	if err != nil && err != sql.ErrNoRows {
		log.Info("Create account error", zap.Error(err))
		return err
	}

	return nil
}

func (r *dBOps) GetAccount(ctx context.Context, id int64) (*models.Account, error) {

	var account models.Account

	if err := r.db.Get(&account, getAccountQuery, id); err != nil {
		log.Info("fetch account error", zap.Error(err))
		return nil, err
	}
	return &account, nil
}

func (r *dBOps) GetAccountByDocument(ctx context.Context, docNo string) (*models.Account, error) {

	var account models.Account

	if err := r.db.Get(&account, getAccountByDocumentQuery, docNo); err != nil && err != sql.ErrNoRows {
		log.Info("fetch account bu document error", zap.Error(err))
		return nil, err
	}
	return &account, nil
}
