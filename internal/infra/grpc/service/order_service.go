package service

import (
	"context"

	"github.com/jamersom/clean-arch-order-list/internal/infra/grpc/pb"
	"github.com/jamersom/clean-arch-order-list/internal/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase  usecase.CreateOrderUseCase
	ListOrdersUseCase   usecase.ListOrderUseCase
	GetOrderByIDUseCase usecase.GetOrderByIDUseCase
	UpdateOrderUseCase  usecase.UpdateOrderUseCase
	DeleteOrderUseCase  usecase.DeleteOrderUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	listOrdersUseCase usecase.ListOrderUseCase,
	getOrderIdUseCase usecase.GetOrderByIDUseCase,
	updateOrderUseCase usecase.UpdateOrderUseCase,
	deleteOrderUseCase usecase.DeleteOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase:  createOrderUseCase,
		ListOrdersUseCase:   listOrdersUseCase,
		GetOrderByIDUseCase: getOrderIdUseCase,
		UpdateOrderUseCase:  updateOrderUseCase,
		DeleteOrderUseCase:  deleteOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	output, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var orders []*pb.CreateOrderResponse
	for _, order := range output {
		orders = append(orders, &pb.CreateOrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return &pb.ListOrdersResponse{
		Orders: orders,
	}, nil
}

func (s *OrderService) GetOrderById(ctx context.Context, in *pb.GetOrderByIdRequest) (*pb.GetOrderByIdResponse, error) {
	dto := usecase.OrderInputDTO{
		ID: in.Id,
	}
	output, err := s.GetOrderByIDUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	orderRes := pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}

	return &pb.GetOrderByIdResponse{
		Order: &orderRes,
	}, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, in *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.UpdateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, in *pb.DeleteOrderRequest) (*emptypb.Empty, error) {
	id := usecase.OrderInputDTO{
		ID: in.Id,
	}

	if err := s.DeleteOrderUseCase.Execute(id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
