package repository

import "github.com/jmoiron/sqlx"

type dBOps struct {
	db *sqlx.DB
}

func NewDBOpsRepository(db *sqlx.DB) *dBOps {
	return &dBOps{db: db}
}
