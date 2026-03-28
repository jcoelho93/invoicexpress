package invoicexpress

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ClientsService struct {
	IxClient *InvoiceXpressClient
	Client   *http.Client
}

func (api *ClientsService) Get(clientId int) (Client, error) {
	endpoint := fmt.Sprintf("%s/%s/%d.json", api.IxClient.Host, "clients", clientId)

	u, err := url.Parse(endpoint)
	if err != nil {
		return Client{}, err
	}

	params := url.Values{}
	params.Add("api_key", api.IxClient.ApiKey)

	u.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return Client{}, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := api.Client.Do(req)
	if err != nil {
		return Client{}, err
	}
	defer resp.Body.Close()

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
