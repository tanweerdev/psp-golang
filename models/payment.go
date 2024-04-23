package models

type Payment struct {
	CardNumber string  `json:"card_number"`
	ExpiryDate string  `json:"expiry_date"`
	CVV        string  `json:"cvv"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	MerchantID string  `json:"merchant_id"`
}
