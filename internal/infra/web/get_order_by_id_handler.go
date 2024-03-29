package web

import (
	"encoding/json"
	"net/http"

	"github.com/jamersom/clean-arch-order-list/internal/entity"
	"github.com/jamersom/clean-arch-order-list/internal/usecase"
	"github.com/jamersom/clean-arch-order-list/pkg/events"
)

type WebGetOrderByIDHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	GetOrderByIDEvent events.EventInterface
}

func NewWebGetOrderByIDHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	GetOrderByIDEvent events.EventInterface,
) *WebGetOrderByIDHandler {
	return &WebGetOrderByIDHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		GetOrderByIDEvent: GetOrderByIDEvent,
	}
}

func (h *WebGetOrderByIDHandler) GetOrderByID(res http.ResponseWriter, req *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	orderUseCase := usecase.NewGetOrderByIDUseCase(h.OrderRepository, h.GetOrderByIDEvent, h.EventDispatcher)
	output, err := orderUseCase.Execute(dto)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(res).Encode(output)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
