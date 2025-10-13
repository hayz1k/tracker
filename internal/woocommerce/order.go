package woocommerce

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"orderTracker/configs"
	"orderTracker/internal/domain/order"
	"orderTracker/internal/service"
	"time"
)

const PROCESSING = "processing"

func UpdateOrders(ctx context.Context, cfg *configs.Config) {

	apiURL := fmt.Sprintf("%s/wp-json/wc/v3/orders?", cfg.WooCommerce.Url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Fatal()
	}

	req.SetBasicAuth(cfg.WooCommerce.Key, cfg.WooCommerce.Secret)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal()
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var orders []order.Order

	var rawOrders []struct {
		ID          int    `json:"id"`
		Status      string `json:"status"`
		DateCreated string `json:"date_created"`
		Shipping    struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Address1  string `json:"address_1"`
			City      string `json:"city"`
			Postcode  string `json:"postcode"`
			Country   string `json:"country"`
		} `json:"shipping"`
	}

	if err := json.Unmarshal(body, &rawOrders); err != nil {
		panic(err)
	}

	for _, raw := range rawOrders {
		created, _ := time.Parse("2006-01-02T15:04:05", raw.DateCreated)

		order := order.Order{
			OrderID:    raw.ID,
			FirstName:  raw.Shipping.FirstName,
			SecondName: raw.Shipping.LastName,
			DeliveryAddress: fmt.Sprintf("%s, %s, %s, %s",
				raw.Shipping.Address1,
				raw.Shipping.City,
				raw.Shipping.Postcode,
				raw.Shipping.Country),
			CurrentStatus: raw.Status,
			Created:       created,
			TrackNumber:   "", // генерируем
		}
		orders = append(orders, order)
	}

	for _, order := range orders {
		if order.CurrentStatus == PROCESSING {
			fmt.Println(order)
		}
		tn, err := service.GenerateTrackNumber(ctx, &order)
		if err != nil {
			log.Err(err).Msg("error to generate track number")
		}
		order.TrackNumber = tn
	}
}
