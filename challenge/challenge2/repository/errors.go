package repository

import (
	"errors"

	"github.com/lib/pq"
)

var (
	ErrNotFound = errors.New("not found")
	ErrDuplicate = errors.New("duplicate")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrInUse = errors.New("in use")
)

// isPgCode untuk check error PostgreSQL dengan kode tertentu
func isPgCode(err error, code string) bool {
	var pqErr *pq.Error
	return errors.As(err, &pqErr) && string(pqErr.Code) == code
}