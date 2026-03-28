package invoicexpress

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type InvoicesService struct {
	IxClient *InvoiceXpressClient
	Client   *http.Client
}

func (api *InvoicesService) Get(documentId int) (Invoice, error) {
	endpoint := fmt.Sprintf("%s/%d.json", "invoices", documentId)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := api.IxClient.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return Invoice{}, err
	}

	resp, err := api.IxClient.Do(req)
	if err != nil {
		return Invoice{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return Invoice{}, fmt.Errorf("HTTP %d", resp.StatusCode)
		}
		apiErr.StatusCode = resp.StatusCode
		return Invoice{}, &apiErr
	}

	body, _ := io.ReadAll(resp.Body)

	var response CreateInvoiceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Invoice{}, err
	}

	return response.Invoice, nil
}

func (api *InvoicesService) Create(request CreateInvoiceRequest) (Invoice, error) {
	endpoint := fmt.Sprintf("%s/%s.json", api.IxClient.Host, "invoices")

	reqBody, err := json.Marshal(request)
	if err != nil {
		return Invoice{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := api.IxClient.NewRequestWithContext(ctx, "POST", endpoint, reqBody)
	if err != nil {
		return Invoice{}, err
	}

	resp, err := api.IxClient.Do(req)
	if err != nil {
		return Invoice{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// greater or equal to 400
	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return Invoice{}, fmt.Errorf("HTTP %d", resp.StatusCode)
		}
		return Invoice{}, &apiErr
	}

	var response CreateInvoiceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Invoice{}, err
	}

	return response.Invoice, nil
}

func (api *InvoicesService) GetItem(itemId int) (InvoiceItem, error) {
	endpoint := fmt.Sprintf("%s/%d.json", "items", itemId)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := api.IxClient.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return InvoiceItem{}, err
	}

	resp, err := api.IxClient.Do(req)
	if err != nil {
		return InvoiceItem{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return InvoiceItem{}, fmt.Errorf("HTTP %d", resp.StatusCode)
		}
		apiErr.StatusCode = resp.StatusCode
		return InvoiceItem{}, &apiErr
	}

	body, _ := io.ReadAll(resp.Body)

	var response GetItemResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return InvoiceItem{}, err
	}

	return response.Item, nil
}

type CreateInvoiceRequest struct {
	Invoice BaseInvoice `json:"invoice"`
}

type CreateInvoiceResponse struct {
	Invoice Invoice `json:"invoice"`
}

type GetItemResponse struct {
	Item InvoiceItem `json:"item"`
}

type InvoiceItem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UnitPrice   string `json:"unit_price"`
	Unit        string `json:"unit"`
	Quantity    string `json:"quantity"`
	Tax         struct {
		ID    int     `json:"id"`
		Name  string  `json:"name"`
		Value float64 `json:"value"`
	} `json:"tax"`
	Discount       float64 `json:"discount"`
	Subtotal       float64 `json:"subtotal"`
	TaxAmount      float64 `json:"tax_amount"`
	DiscountAmount float64 `json:"discount_amount"`
	Total          float64 `json:"total"`
}

type MBReference struct {
	Entity    string  `json:"entity"`
	Value     float64 `json:"value"`
	Reference string  `json:"reference"`
}

type BaseInvoice struct {
	Date         string        `json:"date"`
	DueDate      string        `json:"due_date"`
	TaxExemption string        `json:"tax_exemption"`
	SequenceID   int32         `json:"sequence_id"`
	Client       Client        `json:"client"`
	Items        []InvoiceItem `json:"items"`
}

type Invoice struct {
	BaseInvoice
	ID                     int         `json:"id"`
	Status                 string      `json:"status"`
	Archived               bool        `json:"archived"`
	Type                   string      `json:"type"`
	SequenceNumber         string      `json:"sequence_number"`
	InvertedSequenceNumber string      `json:"inverted_sequence_number"`
	Atcud                  string      `json:"atcud"`
	Reference              string      `json:"reference"`
	Observations           string      `json:"observations"`
	Retention              string      `json:"retention"`
	Permalink              string      `json:"permalink"`
	SaftHash               string      `json:"saft_hash"`
	Sum                    float64     `json:"sum"`
	Discount               float64     `json:"discount"`
	BeforeTaxes            float64     `json:"before_taxes"`
	Taxes                  float64     `json:"taxes"`
	Total                  float64     `json:"total"`
	Currency               string      `json:"currency"`
	MbReference            MBReference `json:"mb_reference"`
}
