package updateorders

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/url"
	"orderTracker/internal/domain/order"
	"orderTracker/internal/domain/site"
	"orderTracker/internal/observability"
	"orderTracker/internal/service"
	"sync"
	"time"
)

const PROCESSING = "pending"

type Service struct {
	siteService  service.SiteService
	orderService service.OrderService
	wooClient    *Client
}

func NewService(siteSvc service.SiteService, orderSvc service.OrderService, woo *Client) *Service {
	return &Service{
		siteService:  siteSvc,
		orderService: orderSvc,
		wooClient:    woo,
	}
}

// UpdateOrders - перебор всех сайтов
func (s *Service) UpdateOrders(ctx context.Context) error {
	observability.OrdersTotal.Inc()

	const workers = 5
	var wg sync.WaitGroup

	sites, err := s.siteService.SiteList(ctx)
	if err != nil {
		return fmt.Errorf("failed to load sites: %w", err)
	}

	if len(sites) == 0 {
		log.Warn().Msg("no sites in database")
		return nil
	}

	sitesCh := make(chan *site.Site, len(sites))
	for _, site := range sites {
		sitesCh <- site
	}
	close(sitesCh)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for site := range sitesCh {
				if err := s.processSite(ctx, site); err != nil {
					log.Error().Err(err).Msgf("worker %d failed to process site %s", workerID, site.Domain)
				}
			}
		}(i)
	}
	wg.Wait()

	return nil
}

// processSite - обработка сайта
func (s *Service) processSite(ctx context.Context, site *site.Site) error {
	rawOrders, err := s.wooClient.GetOrders(ctx, site)
	if err != nil {
		return fmt.Errorf("failed to get orders: %w", err)
	}

	for _, raw := range rawOrders {

		ord := s.buildOrder(raw)

		if exists, _ := s.orderService.Exists(ctx, ord.OrderID); exists {
			log.Info().Msgf("order %d already in database", ord.OrderID)
			continue
		}

		domain := s.extractDomain(raw)
		ord.Site, err = s.siteService.GetSiteByDomain(ctx, domain)
		fmt.Println(ord.Site)
		if err != nil {
			log.Error().Err(err).Msgf("can't resolve site for domain %s", domain)
			continue
		}
		ord.SiteID = ord.Site.ID

		ord.GenerateTrackNumber()
		ord.SetStatusChain()
		fmt.Printf("track number %s, status chain %d", ord.TrackNumber, ord.StatusChain)

		if err := s.orderService.Save(ctx, ord); err != nil {
			log.Error().Err(err).Msgf("can't create order %d", ord.OrderID)
		}
	}
	return nil
}

// buildOrder- конструктор Order из RawOrder
func (s *Service) buildOrder(raw RawOrder) *order.Order {
	created, _ := time.Parse(time.RFC3339, raw.DateCreated)
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
func (s *Service) extractDomain(raw RawOrder) string {

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
