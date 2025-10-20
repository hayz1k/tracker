package dto

import (
	"orderTracker/internal/domain/order"
	"time"
)

type PostOrderRequest struct {
	OrderID         int       `json:"order_id"`
	FirstName       string    `json:"first_name"`
	SecondName      string    `json:"last_name"`
	DeliveryAddress string    `json:"address"`
	Total           string    `json:"total"`
	Created         time.Time `json:"created"`

	SiteID int `json:"site_id"`
}

type CountResponse struct {
	Count int `json:"count"`
}

type OrdersResponse struct {
	Items []*order.Order `json:"items"`
	Total int            `json:"totalCount"`
}
