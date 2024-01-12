package web

import (
	"encoding/json"
	"net/http"

	"github.com/jamersom/clean-arch-order-list/internal/entity"
	"github.com/jamersom/clean-arch-order-list/internal/usecase"
	"github.com/jamersom/clean-arch-order-list/pkg/events"
)

type WebUpdateOrderHandler struct {
	EventDispatcher events.EventDispatcherInterface
	OrderRepository entity.OrderRepositoryInterface
	OrderUpdate     events.EventInterface
}

func NewWebUpdateOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderUpdate events.EventInterface,
) *WebUpdateOrderHandler {
	return &WebUpdateOrderHandler{
		EventDispatcher: EventDispatcher,
		OrderRepository: OrderRepository,
		OrderUpdate:     OrderUpdate,
	}
}

func (h *WebUpdateOrderHandler) UpdateOrder(res http.ResponseWriter, req *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	orderUseCase := usecase.NewUpdateOrderUseCase(h.OrderRepository, h.OrderUpdate, h.EventDispatcher)
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
