package status

import (
	"orderTracker/internal/service"
)

type Handler struct {
	Service service.StatusService
}

func NewStatusHandler(svc service.StatusService) *Handler {
	return &Handler{Service: svc}
}
