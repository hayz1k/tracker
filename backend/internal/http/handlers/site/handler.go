package site

import "orderTracker/internal/service"

type Handler struct {
	Service service.SiteService
}

func NewSiteHandler(s service.SiteService) *Handler {
	return &Handler{Service: s}
}
