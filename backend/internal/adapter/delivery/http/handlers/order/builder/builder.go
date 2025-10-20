package builder

import (
	"orderTracker/internal/adapter/delivery/http/handlers/order/dto"
	"orderTracker/internal/domain/order"
	"orderTracker/internal/domain/site"
)

type Builder struct{}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) BuildOrder(req *dto.PostOrderRequest, site *site.Site) *order.Order {
	ord := &order.Order{
		OrderID:         req.OrderID,
		FirstName:       req.FirstName,
		SecondName:      req.SecondName,
		DeliveryAddress: req.DeliveryAddress,
		Total:           req.Total,
		Created:         req.Created,
		SiteID:          req.SiteID,
	}

	ord.Site = site
	ord.FillOrderData()

	return ord
}
