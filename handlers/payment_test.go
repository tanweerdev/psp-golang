package handlers

import (
	"psp/models"
	"testing"
)

func TestProcessPayment(t *testing.T) {
	// Test valid payment
	validPayment := models.Payment{
		CardNumber: "378282246310005",
		ExpiryDate: "12/25",
		CVV:        "123",
		Amount:     100.00,
		Currency:   "USD",
		MerchantID: "merchant123",
	}
	transaction := processPayment(validPayment)
	if transaction.Status != "Pending" {
		t.Errorf("Expected transaction status to be 'Pending', got '%s'", transaction.Status)
	}

	// Test invalid payment
	invalidPayment := models.Payment{
		CardNumber: "1234567890123456",
		ExpiryDate: "12/25",
		CVV:        "123",
		Amount:     100.00,
		Currency:   "USD",
		MerchantID: "merchant123",
	}
	transaction = processPayment(invalidPayment)
	if transaction.Status != "Invalid Card Number" {
		t.Errorf("Expected transaction status to be 'Invalid Card Number', got '%s'", transaction.Status)
	}
}
