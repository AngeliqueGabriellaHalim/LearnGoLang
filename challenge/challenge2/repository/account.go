package repository

import (
	"challenge2/model"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const accountCols = `id, account_number, account_holder, balance, created_at, updated_at`

// mapErr turns db no row found error into our error
func mapErr(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}
	return err
}

// AccountRepo simpan semua sql related to accounts table
type AccountRepo struct{ db *sqlx.DB }

func (r *AccountRepo) Transfer(context context.Context, fromID uuid.UUID, toID uuid.UUID, amount int64) (any, any) {
	panic("unimplemented")
}

func (r *AccountRepo) ListByAccount(context context.Context, id uuid.UUID) (any, any) {
	panic("unimplemented")
}

// NewAccountRepo constructor
func NewAccountRepo(db *sqlx.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

// Create: insert a new account and return it
func (r *AccountRepo) Create(ctx context.Context, number, holder string, balance int64) (*model.Account, error) {
	var acc model.Account

	var query string = `INSERT INTO accounts (id, account_number, account_holder, balance) VALUES ($1, $2, $3, $4)
	RETURNING ` + accountCols

	//retrieve single row of account thats been created
	err := r.db.GetContext(ctx, &acc, query, uuid.New(), number, holder, balance)

	if isPgCode(err, "23505") {
		return nil, ErrDuplicate
	}

	if err != nil {
		return nil, err
	}

	return &acc, nil
}

// GetAll return semua account, sort by newest creation
func (r *AccountRepo) GetAll(ctx context.Context) ([]model.Account, error) {
	accounts := []model.Account{}

	var query string = `SELECT ` + accountCols + ` FROM accounts
	ORDER BY created_at DESC`

	//return many rows, put into slice
	err := r.db.SelectContext(ctx, &accounts, query)

	return accounts, err
}

func (r *AccountRepo) GetByID(ctx context.Context, id uuid.UUID) (*model.Account, error) {
	var a model.Account

	var query string = `SELECT ` + accountCols + ` FROM accounts WHERE id = $1`

	if err := r.db.GetContext(ctx, &a,
		query, id); err != nil {
		return nil, mapErr(err)
	}
	return &a, nil
}

func (r *AccountRepo) Update(ctx context.Context, id uuid.UUID, holder string, balance int64) (*model.Account, error) {
	var a model.Account

	if err := r.db.GetContext(ctx, &a,
		`UPDATE accounts SET account_holder = $1, balance = $2, updated_at = now()
		 WHERE id = $3 RETURNING `+accountCols, holder, balance, id); err != nil {
		return nil, mapErr(err)
	}
	return &a, nil
}

func (r *AccountRepo) Delete(ctx context.Context, id uuid.UUID) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM accounts WHERE id = $1`, id)
	if isPgCode(err, "23503") {
		return ErrInUse
	}
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return ErrNotFound
	}
	return nil
}
