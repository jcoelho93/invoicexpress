package invoicexpress

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
)

type InvoiceXpressClient struct {
	Host        string
	AccountName string
	ApiKey      string
	client      *http.Client

	Invoices *InvoicesService
	Clients  *ClientsService
}

func NewInvoiceXpressClient(accountName string) *InvoiceXpressClient {
	client := &InvoiceXpressClient{
		Host:        fmt.Sprintf("https://%s.app.invoicexpress.com", accountName),
		AccountName: accountName,
		ApiKey:      os.Getenv("INVOICEXPRESS_API_KEY"),
		client:      &http.Client{},
	}

	client.Invoices = &InvoicesService{IxClient: client}
	client.Clients = &ClientsService{IxClient: client}
	return client
}

func (client *InvoiceXpressClient) NewRequestWithContext(ctx context.Context, method, path string, body []byte) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", client.Host, path)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// pass api_key in query params
	q := req.URL.Query()
	q.Add("api_key", client.ApiKey)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (client *InvoiceXpressClient) Do(req *http.Request) (*http.Response, error) {
	return client.client.Do(req)
}
