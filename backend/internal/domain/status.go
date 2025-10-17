package domain

import "time"

const (
	Created         = "Order created."
	Shipped         = "Package shipped by a sender."
	Sorted          = "Sorted in sorting center."
	WaitingDelivery = "Waiting an a courier."
	Received        = "Order received."
)

type StatusList struct {
	OrderID int
	Name    string
	MinTTL  time.Duration
}

type Status struct {
	Current string
	Expires time.Duration
}

func NewStatus() *Status {
	return &Status{Current: Created, Expires: time.Second}
}
