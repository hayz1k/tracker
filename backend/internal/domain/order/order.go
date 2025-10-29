package order

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"orderTracker/internal/domain/site"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

type Order struct {
	ID              int       `json:"db_id"`
	OrderID         int       `json:"order_id"`
	FirstName       string    `json:"first_name"`
	SecondName      string    `json:"last_name"`
	DeliveryAddress string    `json:"address"`
	Total           string    `json:"total"`
	CurrentStatus   string    `json:"status"`
	Created         time.Time `json:"created"`
	TrackNumber     string    `json:"track_number"`

	SiteID int        `json:"site_id"`
	Site   *site.Site `json:"site"`
}

type OrderFilter struct {
	Status string
	SiteID int
	Search string
}

func (o *Order) GenerateTrackNumber() {
	year := time.Now().Format("06")
	randomPart, _ := generateRandomString(8)
	uniquePart := fmt.Sprintf("%06d", o.OrderID%1000000)
	result := year + randomPart + uniquePart
	o.TrackNumber = result
}

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}

func (o *Order) Validate() error {
	// todo: errors
	if o.OrderID == 0 {
		return fmt.Errorf("order ID is required")
	}
	if o.FirstName == "" {
		return fmt.Errorf("first name is required")
	}
	if o.SecondName == "" {
		return fmt.Errorf("second name is required")
	}
	if o.DeliveryAddress == "" {
		return fmt.Errorf("delivery address is required")
	}
	if o.CurrentStatus == "" {
		return fmt.Errorf("status is required")
	}
	return nil
}
