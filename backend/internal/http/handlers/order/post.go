package order

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"orderTracker/internal/domain/order"
)

func (h *Handler) PostOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var ord order.Order
	if err := json.NewDecoder(r.Body).Decode(&ord); err != nil {
		log.Warn().Err(err).Msg("invalid request body")
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	exists, err := h.Service.Exists(ctx, ord.OrderID)

	if exists {
		log.Info().Msgf("order %d already exists", ord.OrderID)
		http.Error(w, "order already exists", http.StatusConflict)
		return
	}

	ord.Site, err = h.SiteService.GetSiteByID(ctx, ord.SiteID)
	if err != nil {
		log.Error().Err(err).Msgf("cannot get site info for order %d", ord.OrderID)
		http.Error(w, "failed to get site info", http.StatusInternalServerError)
		return
	}

	ord.FillOrderData() // генерируем трек, статус и цепочку доставки

	if err := h.Service.Save(ctx, &ord); err != nil {
		log.Error().Err(err).Msgf("failed to save order %d", ord.OrderID)
		http.Error(w, "failed to save order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ord); err != nil {
		log.Error().Err(err).Msgf("failed to encode response for order %d", ord.OrderID)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	log.Info().Msgf("order %d successfully created", ord.OrderID)
}
