# Payment Service Platform (PSP) System

This is a simple Payment Service Platform (PSP) system developed in Go (Golang). It provides endpoints for processing payments and fetching transactions by card number.

## Setup

1. Clone the Repository:

    ```bash
    git clone <repository_url>
    ```

2. Navigate to the Project Directory:

    ```bash
    cd psp-golang
    ```

3. Run the Application:

    ```bash
    go run main.go
    ```

4. Access Endpoints:

- Payment Processing Endpoint: http://localhost:8080/payment (POST)
- Fetch Transactions Endpoint: http://localhost:8080/transactions (POST)
- Check postman collection inside the project for sample requests

## Design Choices

1. Modular Structure:
The project is organized into multiple packages (main, handlers, models, utils) to separate concerns and improve maintainability.
2. HTTP Server:
The application uses Go's built-in HTTP server to handle incoming requests. This eliminates the need for external frameworks and keeps the codebase lightweight.
3. Concurrency:
Goroutines are used for handling communication with the acquirer asynchronously, improving responsiveness by allowing the server to handle multiple requests concurrently.
4. Data Storage:
Transaction records are stored in an in-memory map for simplicity.
5. Validation:
Payment details are validated using Luhn's algorithm to ensure the integrity of card numbers.
6. Testing:
Unit tests are included for major functions (ValidateCardNumber, processPayment)
