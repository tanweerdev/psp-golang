package handlers

import (
	"encoding/json"
	"net/http"
	"psp/models"
	"psp/utils"
	"strconv"
	"sync"
	"time"
)

var (
	transactions = make(map[string]models.Transaction)
	mu           sync.Mutex
)

// handles new payment requests
func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transaction := processPayment(payment)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

// generates a unique transaction ID
func generateID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// processPayment handles payment processing
func processPayment(payment models.Payment) models.Transaction {
	// Lets Validate card number
	if !utils.ValidateCardNumber(payment.CardNumber) {
		return models.Transaction{
			ID:        "",
			Payment:   payment,
			Status:    "Invalid Card Number",
			Timestamp: time.Now(),
		}
	}

	transactionID := generateID()
	transaction := models.Transaction{
		ID:        transactionID,
		Payment:   payment,
		Status:    "Pending",
		Timestamp: time.Now(),
	}

	// Simulate communication with acquirer
	go handleAcquirerResponse(&transaction)

	return transaction
}

// simulates communication with acquirer and updates transaction status
func handleAcquirerResponse(transaction *models.Transaction) {
	// Mock function: approve if last digit of card number is even, deny otherwise
	lastDigit, _ := strconv.Atoi(string(transaction.Payment.CardNumber[len(transaction.Payment.CardNumber)-1]))
	if lastDigit%2 == 0 {
		transaction.Status = "Approved"
	} else {
		transaction.Status = "Denied"
	}

	// Lets Update transaction status
	mu.Lock()
	defer mu.Unlock()
	transactions[transaction.ID] = *transaction
}
