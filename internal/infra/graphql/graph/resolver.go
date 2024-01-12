package graph

import "github.com/jamersom/clean-arch-order-list/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase  usecase.CreateOrderUseCase
	ListOrderUseCase    usecase.ListOrderUseCase
	GetOrderByIDUseCase usecase.GetOrderByIDUseCase
	UpdateOrderUseCase  usecase.UpdateOrderUseCase
	DeleteOrderUseCase  usecase.DeleteOrderUseCase
}
