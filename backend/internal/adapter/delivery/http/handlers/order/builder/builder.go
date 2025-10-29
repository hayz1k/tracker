package builder

import (
	"orderTracker/internal/adapter/delivery/http/handlers/order/dto"
	"orderTracker/internal/domain/order"
	"orderTracker/internal/domain/site"
	"time"
)

const CREATED = "Created"

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
	ord.GenerateTrackNumber()
	ord.CurrentStatus = CREATED
	return ord
}

func (b *Builder) BuildOrderToUpdate(ord *order.Order, patch map[string]interface{}) (*order.Order, error) {
	for key, value := range patch {
		switch key {
		case "first_name":
			if val, ok := value.(string); ok {
				ord.FirstName = val
			}

		case "second_name":
			if val, ok := value.(string); ok {
				ord.SecondName = val
			}

		case "address":
			if val, ok := value.(string); ok {
				ord.DeliveryAddress = val
			}

		case "total":
			if val, ok := value.(string); ok {
				ord.Total = val
			}

		case "created":
			if val, ok := value.(string); ok {
				v, err := time.ParseInLocation("2006-01-02T15:04:05", val, time.UTC)
				if err != nil {
					return nil, err
				}
				ord.Created = v
			}

		}
	}
	return ord, nil
}
