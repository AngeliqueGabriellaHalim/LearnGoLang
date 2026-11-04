package handler

import (
	"challenge2/repository"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	Accounts *repository.AccountRepo
	Transactions *repository.TransactionRepo
}

//Routes register all endpoint
func (h *Handler) Routes() http.Handler{
	mux := http.NewServeMux()
	
	mux.HandleFunc("POST /accounts", h.createAccount)
	mux.HandleFunc("GET /accounts", h.listAccounts)
	mux.HandleFunc("GET /accounts/{id}", h.getAccount)
	mux.HandleFunc("PUT /accounts/{id}", h.updateAccount)
	mux.HandleFunc("DELETE /accounts/{id}", h.deleteAccount)
	mux.HandleFunc("GET /accounts/{id}/transactions", h.listTransactions)
	mux.HandleFunc("POST /transfer", h.transfer)
	return mux
}

func writeError (w http.ResponseWriter, status int, msg string){
	writeJSON (w, status, map[string]string{"error":msg})
}

func writeJSON (w http.ResponseWriter, status int, v any){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// decodeJSON baca body, return true atau false jika gagal + error
func decodeJSON(w http.ResponseWriter, r *http.Request, dst any) bool {
	// r.Body = http.MaxBytesReader(w, r.Body, 1<<20) 
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		if errors.Is(err, io.EOF) {
			writeError(w, http.StatusBadRequest, "request body must not be empty")
		} else {
			writeError(w, http.StatusBadRequest, "invalid JSON body")
		}
		return false
	}
	return true
}

func parseID(w http.ResponseWriter, r *http.Request) (uuid.UUID, bool){
	id, err := uuid.Parse(r.PathValue("id"))
	
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid account id")
		return uuid.Nil, false
	}
	return id, true
}

func handleRepoError (w http.ResponseWriter, err error){
	switch{
	case errors.Is(err, repository.ErrNotFound):
		writeError(w, http.StatusNotFound, "account number is not found")
	case errors.Is(err, repository.ErrDuplicate):
	writeError(w, http.StatusConflict, "account number already exists")
	case errors.Is(err, repository.ErrInsufficientFunds):
		writeError(w, http.StatusUnprocessableEntity, "insufficient balance")
	case errors.Is(err, repository.ErrInUse):
		writeError(w, http.StatusConflict, "account has transactions and cannot be deleted")
	default:
		log.Println("internal error:", err)
		writeError(w, http.StatusInternalServerError, "internal server error")

	}

}