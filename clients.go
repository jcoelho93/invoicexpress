package invoicexpress

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ClientsService struct {
	IxClient *InvoiceXpressClient
	Client   *http.Client
}

func (api *ClientsService) Get(clientId int) (Client, error) {
	endpoint := fmt.Sprintf("%s/%d.json", "clients", clientId)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := api.IxClient.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return Client{}, err
	}

	resp, err := api.IxClient.Do(req)
	if err != nil {
		return Client{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return Client{}, fmt.Errorf("HTTP %d", resp.StatusCode)
		}
		apiErr.StatusCode = resp.StatusCode
		return Client{}, &apiErr
	}

	body, _ := io.ReadAll(resp.Body)

	var response GetClientResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Client{}, err
	}

	return response.Client, nil
}

type Client struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	City         string `json:"city"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	FiscalID     string `json:"fiscal_id"`
	Website      string `json:"website"`
	Phone        string `json:"phone"`
	Fax          string `json:"fax"`
	Observations string `json:"observations"`
}

type GetClientResponse struct {
	Client Client `json:"client"`
}
