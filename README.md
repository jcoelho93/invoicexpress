# invoicexpress (Go)

Unofficial Go SDK for the InvoiceXpress API.

## Disclaimer

This project is **not an official InvoiceXpress SDK** and is not maintained by InvoiceXpress.

Use this library at your own risk. Always validate behavior in your environment before using it in production.

## Installation

```bash
go get github.com/jcoelho93/invoicexpress
```

## Authentication

Set your API key in an environment variable:

```bash
export INVOICEXPRESS_API_KEY="your_api_key_here"
```

## Simple Example

Fetch an invoice by ID:

```go
package main

import (
	"encoding/json"
	"log"

	"github.com/jcoelho93/invoicexpress"
)

func main() {
	// Replace with your InvoiceXpress account name (subdomain)
	client := invoicexpress.NewInvoiceXpressClient("your-account-name")

	invoice, err := client.Invoices.Get(123456)
	if err != nil {
		log.Fatalf("failed to fetch invoice: %v", err)
	}

	out, err := json.MarshalIndent(invoice, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal invoice: %v", err)
	}

	log.Println(string(out))
}
```

## Notes

- The client reads the API key from `INVOICEXPRESS_API_KEY`.
- Network/API errors are returned from service methods.
