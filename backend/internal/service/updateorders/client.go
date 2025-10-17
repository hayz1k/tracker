package updateorders

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"orderTracker/internal/domain/site"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{httpClient: &http.Client{}}
}

type RawOrder struct {
	ID          int    `json:"id"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
	Billing     struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		City      string `json:"city"`
		Postcode  string `json:"postcode"`
		Country   string `json:"country"`
	} `json:"billing"`
	Total    string `json:"total"`
	MetaData []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"meta_data"`
	PaymentURL string `json:"payment_url"`
}

func (c *Client) GetOrders(ctx context.Context, site *site.Site) ([]RawOrder, error) {
	url := fmt.Sprintf("%s/wp-json/wc/v3/orders", site.Domain)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(site.ConsumerKey, site.ConsumerSecret)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var orders []RawOrder
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}
