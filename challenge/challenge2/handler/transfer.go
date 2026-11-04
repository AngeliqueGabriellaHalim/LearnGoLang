package handler

import (
	"net/http"

	"github.com/google/uuid"
)

type transferRequest struct {
	FromAccountID string `json:"from_account_id"`
	ToAccountID   string `json:"to_account_id"`
	Amount        int64  `json:"amount"`
}

// POST /transfer
func (h *Handler) transfer(w http.ResponseWriter, r *http.Request) {
	var req transferRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	fromID, err1 := uuid.Parse(req.FromAccountID)
	toID, err2 := uuid.Parse(req.ToAccountID)
	if err1 != nil || err2 != nil {
		writeError(w, http.StatusBadRequest, "from_account_id and to_account_id must be valid UUIDs")
		return
	}

	
	if req.Amount <= 0 {
		writeError(w, http.StatusBadRequest, "amount must be greater than 0")
		return
	}
	if fromID == toID { 
		writeError(w, http.StatusBadRequest, "cannot transfer to the same account")
		return
	}

	result, err := h.Transactions.Transfer(r.Context(), fromID, toID, req.Amount)
	if err != nil { // ErrNotFound  404, ErrInsufficientFunds 422
		handleRepoError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, result)
}

// GET /accounts/{id}/transactions
func (h *Handler) listTransactions(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	
	if _, err := h.Accounts.GetByID(r.Context(), id); err != nil {
		handleRepoError(w, err)
		return
	}
	list, err := h.Transactions.ListByAccount(r.Context(), id)
	if err != nil {
		handleRepoError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, list)
}