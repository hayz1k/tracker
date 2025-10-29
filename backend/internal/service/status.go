package service

import (
	"context"
	"orderTracker/internal/domain/status"
)

type StatusService interface {
	SaveStatus(ctx context.Context, status *status.Status) error
	GetByID(ctx context.Context, id int) ([]*status.Status, error)
	GetByData(ctx context.Context, orderID, siteID int) ([]*status.Status, error)
}

type statusService struct {
	store status.Store
}

func NewStatusService(store status.Store) StatusService {
	return &statusService{
		store: store,
	}
}

func (s *statusService) SaveStatus(ctx context.Context, status *status.Status) error {
	return s.store.Save(ctx, status)
}

func (s *statusService) GetByID(ctx context.Context, id int) ([]*status.Status, error) {
	return s.store.GetByID(ctx, id)
}

func (s *statusService) GetByData(ctx context.Context, orderID, siteID int) ([]*status.Status, error) {
	return s.store.GetByData(ctx, orderID, siteID)
}
