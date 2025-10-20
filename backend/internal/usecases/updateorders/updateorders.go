package updateorders

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"orderTracker/internal/domain/site"
	"orderTracker/internal/infrastructure/httpclient/woocommerce"
	"orderTracker/internal/service"
	"sync"
)

const PROCESSING = "pending"

type Service struct {
	siteService  service.SiteService
	orderService service.OrderService
	wooClient    *woocommerce.Client
}

func NewService(siteSvc service.SiteService, orderSvc service.OrderService, woo *woocommerce.Client) *Service {
	return &Service{
		siteService:  siteSvc,
		orderService: orderSvc,
		wooClient:    woo,
	}
}

// UpdateOrders - перебор всех сайтов
func (s *Service) UpdateOrders(ctx context.Context) error {

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
