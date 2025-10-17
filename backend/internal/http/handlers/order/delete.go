package order

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Error().Err(err).Msg("error converting id to int")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	exists, err := h.Service.Exists(ctx, id)
	if !exists {
		log.Error().Msgf("order %d is not exists", id)
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	deleted, err := h.Service.Delete(ctx, id)
	if err != nil {
		log.Error().Err(err).Msgf("failed to delete order %d", id)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	if !deleted {
		log.Info().Msg("order not found")
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

}
