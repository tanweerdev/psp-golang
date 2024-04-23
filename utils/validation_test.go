package utils

import (
	"testing"
)

func TestValidateCardNumber(t *testing.T) {
	// Test valid card number
	validCardNumber := "30569309025904"
	if !ValidateCardNumber(validCardNumber) {
		t.Errorf("Expected %s to be a valid card number, but it was not", validCardNumber)
	}

	// Test invalid card number
	invalidCardNumber := "1234567890123456"
	if ValidateCardNumber(invalidCardNumber) {
		t.Errorf("Expected %s to be an invalid card number, but it was valid", invalidCardNumber)
	}
}
