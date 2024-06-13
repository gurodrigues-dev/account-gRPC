package repository

import (
	"context"
	"database/sql"

	"github.com/gurodrigues-dev/account-gRPC/types"
)

type IAccountRepository interface {
	CreateBill(ctx context.Context, bill *types.Bill) error
	DeleteBill(ctx context.Context, id *string) error
}

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (a *AccountRepository) CreateBill(ctx context.Context, bill *types.Bill) error {
	sqlQuery := `INSERT INTO bills (name, description, amount, installment) VALUES ($1, $2, $3, $4)`
	_, err := a.db.Exec(sqlQuery, bill.Name, bill.Description, bill.Amount, bill.Installment)
	return err
}

func (a *AccountRepository) DeleteBill(ctx context.Context, id *string) error {
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = tx.Exec("DELETE FROM bills WHERE id = $1", *id)
	return err
}
