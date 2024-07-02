package usecases

import (
	"github.com/pitchter/orderRealtime/internal/adapters/repositories"
	"github.com/pitchter/orderRealtime/internal/entities"
	// "github.com/pitchter/orderRealtime/internal/repositories"
)

type OrderUsecase struct {
    orderRepo repositories.OrderRepository
}

func NewOrderUsecase(repo repositories.OrderRepository) *OrderUsecase {
    return &OrderUsecase{orderRepo: repo}
}

func (uc *OrderUsecase) CreateOrder(order entities.Order) (entities.Order, error) {
    return uc.orderRepo.CreateOrder(order)
}
