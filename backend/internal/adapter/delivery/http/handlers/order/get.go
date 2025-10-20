package order

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"orderTracker/internal/adapter/delivery/http/handlers/order/dto"
	"orderTracker/internal/domain/order"
	"strconv"
)

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 0 {
		limit = 10
	}

	filter := order.OrderFilter{
		Status: r.URL.Query().Get("status"),
		Search: r.URL.Query().Get("search"),
	}
	if siteIDStr := r.URL.Query().Get("site_id"); siteIDStr != "" {
		if siteID, err := strconv.Atoi(siteIDStr); err == nil {
			filter.SiteID = siteID
		}
	}

	log.Info().
		Int("page", page).
		Int("limit", limit).
		Str("status", filter.Status).
		Str("search", filter.Search).
		Int("site_id", filter.SiteID).
		Msg("fetching orders")

	orders, err := h.Service.ListOrders(ctx, page, limit, &filter)
	if err != nil {
		log.Error().Err(err).Msg("failed to get orders")
		http.Error(w, "failed to get orders", http.StatusInternalServerError)
		return
	}

	for _, order := range orders {
		order.Site, err = h.SiteService.GetSiteByID(ctx, order.SiteID)
		if err != nil {
			log.Error().Err(err).Int("site_id", order.SiteID).Msg("failed to get site info")
			http.Error(w, "failed to get site info", http.StatusInternalServerError)
			return
		}
	}

	total, err := h.Service.Count(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get orders count")
		http.Error(w, "failed to get orders count", http.StatusInternalServerError)
		return
	}

	resp := dto.OrdersResponse{
		Items: orders,
		Total: total,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
