package repository

import (
	"context"
	"transaction-service/models"
)

func (r *dBOps) CreateAccount(ctx context.Context, profile *models.Account) (*models.Account, error) {

	/*var id int64
	err := r.db.QueryRowContext(
		ctx,
		createOTPQuery,
		otp.ProfileID,
		otp.OTP,
		otp.Validated,
		otp.CreatedAt,
		otp.Expiry,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	otp.ID = id*/
	return nil, nil
}

func (r *dBOps) GetAccount(ctx context.Context, id int64) (*models.Account, error) {

	/*var id int64
	err := r.db.QueryRowContext(
		ctx,
		createOTPQuery,
		otp.ProfileID,
		otp.OTP,
		otp.Validated,
		otp.CreatedAt,
		otp.Expiry,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	otp.ID = id*/
	return nil, nil
}
