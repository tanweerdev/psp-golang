package handlers

import (
	"encoding/json"
	"net/http"
	"psp/models"
)

type TransactionsRequest struct {
	CardNumber string `json:"card_number"`
}

// handles requests to fetch transactions by card number
func TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	var req TransactionsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Please provide card number in request", http.StatusBadRequest)
		return
	}

	var result []models.Transaction
	mu.Lock()
	defer mu.Unlock()
	for _, transaction := range transactions {
		if transaction.Payment.CardNumber == req.CardNumber {
			result = append(result, transaction)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
