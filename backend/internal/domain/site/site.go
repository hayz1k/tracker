package site

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"strings"
)

type Site struct {
	ID             int    `json:"site_id"`
	Domain         string `json:"domain"`
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	Note           string `json:"note"`
	Merchant       string `json:"merchant"`
}

func NewSite(id int, domain, consumerKey, consumerSecret, note, merchant string) (*Site, error) {
	if domain == "" {
		return nil, errors.New("domain is required")
	}
	if consumerKey == "" {
		return nil, errors.New("consumerKey is required")
	}
	if consumerSecret == "" {
		return nil, errors.New("consumerSecret is required")
	}
	if note == "" {
		log.Info().Msg("new site without note")
	}
	if merchant == "" {
		log.Info().Msg("new site without merchant")
	}

	return &Site{
		ID:             id,
		Domain:         domain,
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		Note:           note,
		Merchant:       merchant,
	}, nil
}

func (s *Site) Validate() error {
	if s.Domain == "" || strings.HasPrefix(s.Domain, "http://") {
		fmt.Print("yes0")
		return errors.New("domain is required")
	}

	if !strings.HasPrefix(s.Domain, "https://") {
		s.Domain = "https://" + s.Domain
	}

	if s.ConsumerKey == "" {
		fmt.Print("yes")
		return errors.New("consumerKey is required")
	}
	if s.ConsumerSecret == "" {
		fmt.Print("yes3")
		return errors.New("consumerSecret is required")
	}
	return nil
}
