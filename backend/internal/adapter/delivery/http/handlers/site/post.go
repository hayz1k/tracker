package site

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"orderTracker/internal/domain/site"
)

func (h *Handler) PostSite(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("start posting site")
	ctx := r.Context()

	var s site.Site
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		log.Error().Err(err).Msg("cannot decode request body")
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if valid := s.Validate(); valid != nil {
		log.Info().Msgf("decoded site is not valid: %s", s.Domain)
		http.Error(w, "site is not valid", http.StatusBadRequest)
		return
	}

	if exists, _ := h.Service.SiteExists(ctx, s.Domain); exists {
		log.Info().Msgf("site already exists: %s", s.Domain)
		http.Error(w, "site already exists", http.StatusConflict)
		return
	}

	err := h.Service.SaveSite(ctx, &s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		log.Error().Err(err).Msg("failed to write response")
	}
}
