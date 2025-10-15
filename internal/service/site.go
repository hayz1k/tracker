package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"orderTracker/internal/domain/site"
)

type SiteService interface {
	SaveSite(ctx context.Context, site *site.Site) error
	SiteList(ctx context.Context) ([]*site.Site, error)
	SiteExists(ctx context.Context, domain string) (bool, error)
	GetSiteByID(ctx context.Context, id int) (*site.Site, error)
	GetSiteByDomain(ctx context.Context, domain string) (*site.Site, error)
}

type siteService struct {
	store site.Store
}

func NewSiteService(store site.Store) SiteService {
	return &siteService{
		store: store,
	}
}

func (s *siteService) SaveSite(ctx context.Context, site *site.Site) error {
	log.Info().Msg("start saving site")
	err := s.store.Save(ctx, site)
	if err != nil {
		log.Error().Msg("error saving site")
		return err
	}
	log.Info().Msg("success saving site")
	return nil
}
func (s *siteService) SiteList(ctx context.Context) ([]*site.Site, error) {
	return s.store.FindAll(ctx)
}

func (s *siteService) GetSiteByID(ctx context.Context, id int) (*site.Site, error) {
	return s.store.GetByID(ctx, id)
}

func (s *siteService) GetSiteByDomain(ctx context.Context, domain string) (*site.Site, error) {
	return s.store.GetByDomain(ctx, domain)
}

func (s *siteService) SiteExists(ctx context.Context, domain string) (bool, error) {
	return s.store.Exists(ctx, domain)
}
