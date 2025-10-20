package site

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetSites(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sites, err := h.Service.SiteList(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sites)
}
