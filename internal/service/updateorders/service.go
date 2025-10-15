package updateorders

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/url"
	"orderTracker/internal/domain/order"
	"orderTracker/internal/service"
	"orderTracker/internal/woocommerce"
	"time"
)

type Service struct {
	siteService  service.SiteService
	orderService service.OrderService
	wooClient    *woocommerce.Client // общий HTTP клиент без ключей
}

func NewService(siteSvc service.SiteService, orderSvc service.OrderService, woo *woocommerce.Client) *Service {
	return &Service{
		siteService:  siteSvc,
		orderService: orderSvc,
		wooClient:    woo,
	}
}

func (s *Service) UpdateOrders(ctx context.Context) error {
	sites, err := s.siteService.SiteList(ctx)
	if err != nil {
		return fmt.Errorf("failed to load sites: %w", err)
	}
	if len(sites) == 0 {
		log.Warn().Msg("no sites in database")
		return errors.New("sites not found")
	}
	for _, site := range sites {
		rawOrders, err := s.wooClient.GetOrders(ctx, site)
		if err != nil {
			return fmt.Errorf("failed to get orders for site %s: %w", site.Domain, err)
		}

		for _, raw := range rawOrders {
			if raw.Status != "on-hold" {
				continue
			}

			created, _ := time.Parse(time.RFC3339, raw.DateCreated)

			ord := order.Order{raw.ID,
				raw.Billing.FirstName,
				raw.Billing.LastName,
				fmt.Sprintf("%s, %s, %s, %s",
					raw.Billing.Address1,
					raw.Billing.City,
					raw.Billing.Postcode,
					raw.Billing.Country),
				raw.Total,
				raw.Status,
				created, "", -1, nil, nil, -1}

			if _, err = s.orderService.GetOrderByID(ctx, ord.OrderID); err == nil {
				log.Error().Err(err).Msg("order already in database")
				continue
			}

			ord.GenerateTrackNumber()

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
			domain := u.Host
			ord.Site, err = s.siteService.GetSiteByDomain(ctx, domain)

			if err != nil {
				fmt.Println(domain)
				log.Error().Err(err).Msg("can't get site by domain while updating orders")
				continue
			}

			if err := s.orderService.SaveOrder(ctx, &ord); err != nil {
				log.Error().Err(err).Msg("can't save order while updating orders")
				return err
			}
		}
	}

	return nil
}
