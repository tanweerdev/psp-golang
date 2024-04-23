package main

import (
	"fmt"
	"log"
	"net/http"

	"psp/handlers"
)

func main() {
	http.HandleFunc("/payment", handlers.PaymentHandler)

	http.HandleFunc("/transactions", handlers.TransactionsHandler)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
