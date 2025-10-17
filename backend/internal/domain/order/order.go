package order

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"math/rand"
	"orderTracker/internal/domain/site"
	"orderTracker/internal/domain/status"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

type Order struct {
	OrderID         int       `json:"order_id"`
	FirstName       string    `json:"first_name"`
	SecondName      string    `json:"last_name"`
	DeliveryAddress string    `json:"address"`
	Total           string    `json:"total"`
	CurrentStatus   string    `json:"status"`
	Created         time.Time `json:"created"`
	TrackNumber     string    `json:"track_number"`

	// Сайт (domain, key, secret, note(sender))
	SiteID int        `json:"site_id"`
	Site   *site.Site `json:"site"`
	// Статусы
	StatusChain []string
	StatusIndex int
	IsCustom    bool
}

type OrderFilter struct {
	Status string
	SiteID int
	Search string
}

func StartTTLChecker() {

}

func (o *Order) FillOrderData() {
	o.GenerateTrackNumber()
	o.SetStatusChain()
}

func (o *Order) NextStatus() (string, error) {
	if o.IsCustom {
		log.Info().Msgf("order %d: status is custom", o.OrderID)
		return "", errors.New("status is custom")
	}

	if o.StatusChain == nil {
		o.SetStatusChain()
	}

	if o.StatusIndex >= len(o.StatusChain)-1 {
		return "", errors.New("status chain finished")
	}
	o.StatusIndex++
	o.CurrentStatus = o.StatusChain[o.StatusIndex]
	return o.CurrentStatus, nil
}

func (o *Order) SetStatusChain() {
	o.StatusChain = status.Chains[r.Intn(len(status.Chains))]
	o.StatusIndex = 0
	o.CurrentStatus = o.StatusChain[o.StatusIndex]
	o.IsCustom = false
}

func (o *Order) SetCustomStatus(status string) {
	o.CurrentStatus = status
	o.IsCustom = true
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
