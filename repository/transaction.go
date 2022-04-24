package repository

import (
	"context"
	"transaction-service/models"
)

func (r *dBOps) CreateTransaction(ctx context.Context, profile *models.Transaction) error {

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
	return nil
}
