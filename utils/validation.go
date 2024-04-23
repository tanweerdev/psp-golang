package utils

// checks if the card is valid using Luhn's algorithm
func ValidateCardNumber(cardNumber string) bool {
	// copied code. looks like its standard implementation of Luhn's algorithm
	sum := 0
	double := false
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}

	return sum%10 == 0
}
