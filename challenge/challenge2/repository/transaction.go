package repository

import (
	"challenge2/model"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TransactionRepo struct{ db *sqlx.DB }

func NewTransactionRepo(db *sqlx.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}


func (r *TransactionRepo) Transfer(ctx context.Context, fromID, toID uuid.UUID, amount int64) (*model.TransferResult, error) {
	// all step either succedd together or none
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	
	defer tx.Rollback()

	//lock both acc rows
	var locked []struct {
		ID uuid.UUID `db:"id"`
		Balance int64     `db:"balance"`
	}
	if err := tx.SelectContext(ctx, &locked,
		`SELECT id, balance FROM accounts WHERE id IN ($1, $2) ORDER BY id FOR UPDATE`,
		fromID, toID); err != nil {
		return nil, err
	}

	//check both acc exist
	if len(locked) != 2 {
		return nil, ErrNotFound
	}
	
	//balance check after lock
	for _, l := range locked {
		if l.ID == fromID && l.Balance < amount {
			return nil, ErrInsufficientFunds
		}
	}

	// move the money
	var res model.TransferResult
	if err := tx.GetContext(ctx, &res.FromAccount,
		`UPDATE accounts SET balance = balance - $1, updated_at = now()
		 WHERE id = $2 RETURNING `+accountCols, amount, fromID); err != nil {
		return nil, err
	}
	if err := tx.GetContext(ctx, &res.ToAccount,
		`UPDATE accounts SET balance = balance + $1, updated_at = now()
		 WHERE id = $2 RETURNING `+accountCols, amount, toID); err != nil {
		return nil, err
	}

	//record transfer history in transaction table
	if err := tx.GetContext(ctx, &res.Transaction,
		`INSERT INTO transactions (id, from_account_id, to_account_id, amount)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, from_account_id, to_account_id, amount, created_at`,
		uuid.New(), fromID, toID, amount); err != nil {
		return nil, err
	}

	//commit to permanent
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &res, nil
}

// ListByAccount returns every transaction involving a particular account
func (r *TransactionRepo) ListByAccount(ctx context.Context, accountID uuid.UUID) ([]model.Transaction, error) {
	list := []model.Transaction{} 
	err := r.db.SelectContext(ctx, &list,
		`SELECT id, from_account_id, to_account_id, amount, created_at
		 FROM transactions
		 WHERE from_account_id = $1 OR to_account_id = $1
		 ORDER BY created_at DESC`, accountID)
	return list, err
}
