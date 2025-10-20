package updateorders

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) UpdateOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.Service.UpdateOrders(ctx); err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("failed to update orders")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("orders updated successfully"))
}
