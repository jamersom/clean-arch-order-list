package web

import (
	"encoding/json"
	"net/http"

	"github.com/jamersom/clean-arch-order-list/internal/entity"
	"github.com/jamersom/clean-arch-order-list/internal/usecase"
	"github.com/jamersom/clean-arch-order-list/pkg/events"
)

type WebListOrdersHandler struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewWebListOrdersHandler(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *WebListOrdersHandler {
	return &WebListOrdersHandler{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (h *WebListOrdersHandler) List(res http.ResponseWriter, req *http.Request) {
	orderUseCase := usecase.NewListOrderUseCase(h.OrderRepository, h.OrderListed, h.EventDispatcher)
	output, err := orderUseCase.Execute()
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
