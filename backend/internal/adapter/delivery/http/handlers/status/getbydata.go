package status

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func (h *Handler) GetByData(w http.ResponseWriter, r *http.Request) {
	orderIDStr := r.URL.Query().Get("order_id")
	siteIDStr := r.URL.Query().Get("site_id")

	if orderIDStr == "" || siteIDStr == "" {
		http.Error(w, "order_id and site_id are required", http.StatusBadRequest)
		return
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		http.Error(w, "invalid order_id", http.StatusBadRequest)
		return
	}

	siteID, err := strconv.Atoi(siteIDStr)
	if err != nil {
		http.Error(w, "invalid site_id", http.StatusBadRequest)
		return
	}

	statuses, err := h.Service.GetByData(r.Context(), orderID, siteID)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get statuses for order %d, site %d", orderID, siteID)
		http.Error(w, "failed to get statuses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		log.Error().Err(err).Msg("failed to encode response")
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
