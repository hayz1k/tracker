package status

import (
	"errors"
	"time"
)

type Status struct {
	ID          int       `json:"id"`
	OrderID     int       `json:"order_id"`
	SiteID      int       `json:"site_id"`
	Status      string    `json:"status"`
	StatusIndex int       `json:"status_index"`
	CreatedAt   time.Time `json:"created_at"`
	IsCustom    bool      `json:"is_custom"`
}

func (s *Status) SetCustomStatus(status string) {
	s.Status = status
	s.IsCustom = true
}

func (s *Status) SetStatus() {
	s.StatusIndex = 0
	chain := Chains[s.OrderID%len(Chains)]
	s.Status = chain[s.StatusIndex]
	s.CreatedAt = time.Time{}
}

func (s *Status) NextStatus() (*Status, error) {
	if s.IsCustom {
		return nil, errors.New("status is custom")
	}

	chain := Chains[s.OrderID%len(Chains)]
	// s = 4   len = 5(но ласт индекс 4)
	if s.StatusIndex >= len(chain)-1 {
		return nil, errors.New("not enough elements")
	}
	s.StatusIndex++
	s.Status = chain[s.StatusIndex]

	next := &Status{
		OrderID:     s.OrderID,
		SiteID:      s.SiteID,
		Status:      chain[s.StatusIndex+1],
		StatusIndex: s.StatusIndex + 1,
		IsCustom:    false,
		CreatedAt:   time.Now(),
	}

	return next, nil
}
