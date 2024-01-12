package web

import (
	"encoding/json"
	"net/http"

	"github.com/jamersom/clean-arch-order-list/internal/entity"
	"github.com/jamersom/clean-arch-order-list/internal/usecase"
	"github.com/jamersom/clean-arch-order-list/pkg/events"
)

type WebDeleteOrderHandler struct {
	EventDispatcher  events.EventDispatcherInterface
	OrderRepository  entity.OrderRepositoryInterface
	DeleteOrderEvent events.EventInterface
}

func NewWebDeleteOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	DeleteOrderEvent events.EventInterface,
) *WebDeleteOrderHandler {
	return &WebDeleteOrderHandler{
		EventDispatcher:  EventDispatcher,
		OrderRepository:  OrderRepository,
		DeleteOrderEvent: DeleteOrderEvent,
	}
}

func (h *WebDeleteOrderHandler) DeleteOrder(res http.ResponseWriter, req *http.Request) {
	var dto usecase.OrderInputDTO
	if err := json.NewDecoder(req.Body).Decode(&dto); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	orderUseCase := usecase.NewDeleteOrderUseCase(h.OrderRepository, h.DeleteOrderEvent, h.EventDispatcher)
	if err := orderUseCase.Execute(dto); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
