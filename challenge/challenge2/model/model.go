package model

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID uuid.UUID `json:"id" db:"id"`
	BankID *uuid.UUID `json:"bank_id,omitempty" db:"bank_id"`
	AccountNumber string     `json:"account_number" db:"account_number"`
	AccountHolder string `json:"account_holder" db:"account_holder"`
	Balance int64 `json:"balance" db:"balance"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}

type Transaction struct {
	ID uuid.UUID `json:"id" db:"id"`
	FromAccountID uuid.UUID `json:"from_account_id" db:"from_account_id"`
	ToAccountID uuid.UUID `json:"to_account_id" db:"to_account_id"`
	Amount int64 `json:"amount" db:"amount"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type TransferResult struct {
	Transaction Transaction `json:"transaction"`
	FromAccount Account     `json:"from_account"`
	ToAccount Account     `json:"to_account"`
}