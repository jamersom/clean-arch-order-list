package web

import (
	"encoding/json"
	"net/http"

	"github.com/jamersom/clean-arch-order-list/internal/entity"
	"github.com/jamersom/clean-arch-order-list/internal/usecase"
	"github.com/jamersom/clean-arch-order-list/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

func (h *WebOrderHandler) Create(res http.ResponseWriter, req *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	orderUseCase := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
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
