package site

import "context"

type Store interface {
	Save(ctx context.Context, site *Site) error
	GetByID(ctx context.Context, id int) (*Site, error)
	GetByDomain(ctx context.Context, domain string) (*Site, error)
	FindAll(ctx context.Context) ([]*Site, error)
	Exists(ctx context.Context, domain string) (bool, error)
	// Delete(ctx context.Context, id int) error
	// Update(ctx context.Context, site *Site) error
}
