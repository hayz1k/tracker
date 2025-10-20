package updateorders

import (
	"orderTracker/internal/usecases/updateorders"
)

type Handler struct {
	Service *updateorders.Service
}

func NewHandler(svc *updateorders.Service) *Handler {
	return &Handler{Service: svc}
}
