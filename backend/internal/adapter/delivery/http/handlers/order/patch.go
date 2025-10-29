package order

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"orderTracker/internal/adapter/delivery/http/handlers/order/builder"
	"strconv"
)

func (h *Handler) PatchOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var patch map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	order, err := h.Service.GetByID(ctx, id)
	if err != nil {
		http.Error(w, "order not exists", http.StatusNotFound)
		return
	}

	b := builder.NewBuilder()
	order, err = b.BuildOrderToUpdate(order, patch)
	if err != nil {
		http.Error(w, "error building order", http.StatusBadRequest)
		return
	}

	err = h.Service.Update(ctx, order)
	if err != nil {
		http.Error(w, "error while updating order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
