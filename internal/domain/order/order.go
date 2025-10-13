package order

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

type Order struct {
	OrderID         int       `json:"id"`
	FirstName       string    `json:"first_name"`
	SecondName      string    `json:"last_name"`
	DeliveryAddress string    `json:"shipping"`
	CurrentStatus   string    `json:"status"`
	Created         time.Time `json:"date_created"`
	TrackNumber     string    `json:"track_number"`
}

func (o *Order) GenerateTrackNumber() (string, error) {

	year := time.Now().Format("06")

	randomPart, _ := generateRandomString(8)

	uniquePart := fmt.Sprintf("%06d", o.OrderID%1000000)

	result := year + randomPart + uniquePart

	return result, nil
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
