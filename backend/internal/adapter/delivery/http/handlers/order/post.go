package order

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"orderTracker/internal/adapter/delivery/http/handlers/order/builder"
	"orderTracker/internal/adapter/delivery/http/handlers/order/dto"
)

func (h *Handler) PostOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.PostOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Warn().Err(err).Msg("invalid request body")
		h.Metrics.BusinessOrders.WithLabelValues("failed").Inc()
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	exists, err := h.Service.Exists(ctx, req.OrderID)
	if err != nil {
		log.Error().Err(err).Msg("failed to check if order exists")
		h.Metrics.BusinessOrders.WithLabelValues("failed").Inc()
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	if exists {
		log.Info().Msgf("order %d already exists", req.OrderID)
		h.Metrics.BusinessOrders.WithLabelValues("failed").Inc()
		http.Error(w, "order already exists", http.StatusConflict)
		return
	}

	site, err := h.SiteService.GetSiteByID(ctx, req.SiteID)
	if err != nil {
		log.Error().Err(err).Msgf("cannot get site info for order %d", req.OrderID)
		h.Metrics.BusinessOrders.WithLabelValues("failed").Inc()
		http.Error(w, "failed to get site info", http.StatusInternalServerError)
		return
	}

	ord := builder.NewBuilder().BuildOrder(&req, site)

	if err := h.Service.Save(ctx, ord); err != nil {
		log.Error().Err(err).Msgf("failed to save order %d", ord.OrderID)
		h.Metrics.BusinessOrders.WithLabelValues("failed").Inc()
		http.Error(w, "failed to save order", http.StatusInternalServerError)
		return
	}

	h.Metrics.BusinessOrders.WithLabelValues("created").Inc()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	if err := json.NewEncoder(w).Encode(ord); err != nil {
		log.Error().Err(err).Msgf("failed to encode response for order %d", ord.OrderID)
		h.Metrics.BusinessOrders.WithLabelValues("failed").Inc()
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
