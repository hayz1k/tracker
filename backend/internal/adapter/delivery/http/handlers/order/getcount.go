package order

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"orderTracker/internal/adapter/delivery/http/handlers/order/dto"
)

func (h *Handler) GetOrdersCount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	count, err := h.Service.Count(ctx)
	if err != nil {
		log.Error().Err(err).Msg("error getting orders count")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := dto.CountResponse{Count: count}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error().Err(err).Msg("error encoding response")
		http.Error(w, "error encoding response", http.StatusInternalServerError)
		return
	}
	log.Info().Msgf("orders count: %d", count)
}
