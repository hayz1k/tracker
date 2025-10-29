package updateorders

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/url"
	"orderTracker/internal/domain/order"
	"orderTracker/internal/domain/site"
	"orderTracker/internal/infrastructure/httpclient/woocommerce"
	"time"
)

const CREATED = "Created"

// processSite - обработка конкретного сайта
func (s *Service) processSite(ctx context.Context, site *site.Site) error {
	rawOrders, err := s.wooClient.GetOrders(ctx, site)
	if err != nil {
		return fmt.Errorf("failed to get orders: %w", err)
	}

	for _, raw := range rawOrders {

		ord := s.buildOrder(raw)

		if exists, _ := s.orderService.Exists(ctx, ord.OrderID, ord.SiteID); exists {
			continue
		}

		domain := s.extractDomain(raw)
		fmt.Printf("domain %s", domain)
		ord.Site, err = s.siteService.GetSiteByDomain(ctx, domain)
		fmt.Print(ord.Site)
		if err != nil {
			log.Error().Err(err).Msgf("can't resolve site for domain %s", domain)
			continue
		}

		ord.SiteID = ord.Site.ID

		ord.GenerateTrackNumber()
		ord.CurrentStatus = CREATED

		if err := s.orderService.Save(ctx, ord); err != nil {
			log.Error().Err(err).Msgf("can't create order %d", ord.OrderID)
		}
	}
	return nil
}

// buildOrder- конструктор Order из RawOrder
func (s *Service) buildOrder(raw woocommerce.RawOrder) *order.Order {
	created, err := time.ParseInLocation("2006-01-02T15:04:05", raw.DateCreated, time.UTC)
	if err != nil {
		log.Info().Msgf("parse error: %v (input: %v)\n", err, raw.DateCreated)
	}
	return &order.Order{
		OrderID:    raw.ID,
		FirstName:  raw.Billing.FirstName,
		SecondName: raw.Billing.LastName,
		DeliveryAddress: fmt.Sprintf("%s, %s, %s, %s",
			raw.Billing.Address1,
			raw.Billing.City,
			raw.Billing.Postcode,
			raw.Billing.Country),
		Total:         raw.Total,
		CurrentStatus: raw.Status,
		Created:       created,
	}
}

// extractDomain - извлекает домен
func (s *Service) extractDomain(raw woocommerce.RawOrder) string {

	var siteURL string
	for _, meta := range raw.MetaData {
		if meta.Key == "_wc_order_attribution_session_entry" {
			siteURL = meta.Value
			break
		}
	}
	if siteURL == "" {
		siteURL = raw.PaymentURL
	}
	u, _ := url.Parse(siteURL)

	return "https://" + u.Host
}
