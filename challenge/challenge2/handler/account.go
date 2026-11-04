package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

// createAccountRequest: fields that client can send for create account
type createAccountRequest struct{
	AccountNumber string `json:"account_number"`
	AccountHolder string `json:"account_holder"`
	Balance int64 `json:"balance"`
}

// createAccountRequest: fields that client can send for create account
type createAccountResponse struct{
	Message string `json:"message"`
	Data	accountData `json:"data"`
}

type updateAccountRequest struct {
	AccountHolder string `json:"account_holder"`
	Balance       *int64 `json:"balance"`
}

// accountData is create account response's data 
type accountData struct{
	ID uuid.UUID `json:"id"`
	AccountNumber string `json:"account_number"`
	AccountHolder string `json:"account_holder"`
	Balance int64 `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}


//
func (h *Handler) createAccount (w http.ResponseWriter, r *http.Request){
	var req createAccountRequest

	if !decodeJSON(w, r, &req){
		return
	}

	req.AccountNumber = strings.TrimSpace(req.AccountNumber)
	req.AccountHolder = strings.TrimSpace(req.AccountHolder)

	if (req.AccountNumber ==""){
		writeError(w, http.StatusBadRequest, "account_number cannot be empty")
		return
	}
	
	if (req.AccountHolder ==""){
		writeError(w, http.StatusBadRequest, "account_holder cannot be empty")
		return
	}

	if (req.Balance < 0){
		writeError(w, http.StatusBadRequest, "minimum balance is 0")
		return
	}

	acc, err := h.Accounts.Create(r.Context(), req.AccountNumber, req.AccountHolder, req.Balance)
	if err != nil {
		handleRepoError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, createAccountResponse{
		Message: "Account created successfully",
		Data: accountData{
			ID: acc.ID,
			AccountNumber: acc.AccountNumber,
			AccountHolder: acc.AccountHolder,
			Balance: acc.Balance,
			CreatedAt: acc.CreatedAt,
		},
	})
}

// GET /accounts
func (h *Handler) listAccounts(w http.ResponseWriter, r *http.Request) {
	accs, err := h.Accounts.GetAll(r.Context())
	if err != nil {
		handleRepoError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, accs)
}
	
// GET /accounts/{id}
func (h *Handler) getAccount(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r) 
	if !ok {
		return
	}
	acc, err := h.Accounts.GetByID(r.Context(), id) 
	if err != nil {
		handleRepoError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, acc)
}

// PUT /accounts/{id}
func (h *Handler) updateAccount(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	var req updateAccountRequest
	if !decodeJSON(w, r, &req) {
		return
	}
	req.AccountHolder = strings.TrimSpace(req.AccountHolder)
	if req.AccountHolder == "" || req.Balance == nil { 
		writeError(w, http.StatusBadRequest, "account_holder and balance are required")
		return
	}
	if *req.Balance < 0 { 
		writeError(w, http.StatusBadRequest, "balance must not be negative")
		return
	}

	acc, err := h.Accounts.Update(r.Context(), id, req.AccountHolder, *req.Balance)
	if err != nil { 
		handleRepoError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, acc)
}

// DELETE /accounts/{id}
func (h *Handler) deleteAccount(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	if err := h.Accounts.Delete(r.Context(), id); err != nil { // 404 or 409
		handleRepoError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": "Account deleted successfully"})
}