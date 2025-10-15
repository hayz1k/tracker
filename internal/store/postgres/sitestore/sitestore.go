package sitestore

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rs/zerolog/log"
	"orderTracker/internal/domain/site"
)

type SiteStore struct {
	db *sql.DB
}

func NewSiteStore(db *sql.DB) *SiteStore {
	return &SiteStore{db: db}
}

func (s *SiteStore) Save(ctx context.Context, site *site.Site) error {
	log.Info().Msg("saving site")
	query := "INSERT INTO sites " +
		"(domain, consumer_key, consumer_secret, note, merchant) " +
		"VALUES($1,$2,$3,$4,$5)"
	_, err := s.db.Exec(query, site.Domain, site.ConsumerKey, site.ConsumerSecret, site.Note, site.Merchant)
	if err != nil {
		log.Error().Err(err).Msg("error adding order")
		return errors.New("internal error")
	}
	log.Info().Msg("successfully added order")
	return nil
}

func (s *SiteStore) GetByID(ctx context.Context, id int) (*site.Site, error) {
	log.Info().Msg("finding site by id")
	var result site.Site
	query := "SELECT * FROM sites WHERE id = $1"
	err := s.db.QueryRow(query, id).
		Scan(&result.ID, &result.Domain, &result.ConsumerKey,
			&result.ConsumerSecret, &result.Note, &result.Merchant)
	if err != nil {
		log.Error().Err(err).Msg("error getting site by id")
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("site not found")
		}
		return nil, errors.New("internal error")
	}
	log.Info().Msg("success getting site by id")
	return &result, nil
}

func (s *SiteStore) GetByDomain(ctx context.Context, domain string) (*site.Site, error) {
	log.Info().Msg("finding site by domain")
	var result site.Site
	query := "SELECT * FROM sites WHERE domain = $1"
	err := s.db.QueryRow(query, domain).
		Scan(&result.ID, &result.Domain, &result.ConsumerKey,
			&result.ConsumerSecret, &result.Note, &result.Merchant)
	if err != nil {
		log.Error().Err(err).Msg("error getting site by domain")
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("site not found")
		}
		return nil, errors.New("internal error")
	}
	log.Info().Msg("success getting site by domain")
	return &result, nil
}

func (s *SiteStore) FindAll(ctx context.Context) ([]*site.Site, error) {
	log.Info().Msg("finding all sites")
	rows, err := s.db.Query("SELECT domain, consumer_key, consumer_secret, note, merchant FROM sites")
	if err != nil {
		log.Error().Msg("error selecting sites")
		return nil, err
	}
	defer rows.Close()

	var sites []*site.Site

	for rows.Next() {
		var (
			domain         string
			consumerKey    string
			consumerSecret string
			note           string
			merchant       string
		)

		if err := rows.Scan(&domain, &consumerKey, &consumerSecret, &note, &merchant); err != nil {
			return nil, err
		}

		siteEntity, err := site.NewSite(domain, consumerKey, consumerSecret, note, merchant)
		if err != nil {
			return nil, err
		}
		sites = append(sites, siteEntity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Info().Msg("success finding all sites")
	return sites, nil
}

func (s *SiteStore) Exists(ctx context.Context, domain string) (bool, error) {
	log.Info().Msg("start searching site")
	var tmp int
	query := "SELECT 1 FROM sites WHERE domain = $1 LIMIT 1"
	err := s.db.QueryRowContext(ctx, query, domain).Scan(&tmp)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errors.New("site is not exists")
		}
		return false, errors.New("internal error")
	}
	log.Info().Msg("site found")
	return true, nil
}

func (s *SiteStore) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (s *SiteStore) Update(site *site.Site) error {
	//TODO implement me
	panic("implement me")
}
