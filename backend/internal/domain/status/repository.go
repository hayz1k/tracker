package status

import "context"

type Store interface {
	Save(ctx context.Context, status *Status) error
	GetByID(ctx context.Context, id int) ([]*Status, error)
	GetByData(ctx context.Context, orderID, siteID int) ([]*Status, error)
}
