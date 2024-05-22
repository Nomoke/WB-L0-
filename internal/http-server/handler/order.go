package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nomoke/wb-test-app/internal/service"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type Order struct {
	service service.Order
}

func New(service service.Order) *Order {
	return &Order{service: service}
}

func (h *Order) GetOrderByID(w http.ResponseWriter, r *http.Request, log *slog.Logger) {
	id, err := orderIdValidator(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("uuid is invalid: %s", err), http.StatusBadRequest)
		log.Error("uuid is invalid: %s", err)
		return
	}

	order, err := h.service.GetOrderById(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while getting order by ID: %s", err), http.StatusInternalServerError)
		log.Error(fmt.Sprintf("error while getting order by ID: %s", err), err)
		return
	}

	orderJson, err := json.Marshal(order)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while making json: %s", err), http.StatusInternalServerError)
		log.Error(fmt.Sprintf("error while making json: %s", err), err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(orderJson)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while writing json: %s", err), http.StatusInternalServerError)
		log.Error(fmt.Sprintf("error while writing json: %s", err), err)
		return
	}
}

func orderIdValidator(id string) (uuid.UUID, error) {
	orderUid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, err
	}

	return orderUid, nil
}
