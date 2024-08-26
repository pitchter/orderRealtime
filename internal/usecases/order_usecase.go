package usecases

import (
	"errors"

	"github.com/pitchter/orderRealtime/internal/adapters/repositories"
	"github.com/pitchter/orderRealtime/internal/entities"
	services "github.com/pitchter/orderRealtime/internal/service"
)

type OrderUsecase struct {
	orderRepo repositories.OrderRepository
	menuRepo  repositories.MenuRepository
}

func NewOrderUsecase(orderRepo repositories.OrderRepository, menuRepo repositories.MenuRepository) *OrderUsecase {
	return &OrderUsecase{orderRepo: orderRepo, menuRepo: menuRepo}
}

func (uc *OrderUsecase) CreateOrder(order entities.Order) (entities.Order, error) {
	var total float64
	var fullItems []entities.OrderItem

	for _, item := range order.Items {
		menuItem, err := uc.menuRepo.GetMenuItemByID(item.MenuItemID)
		if err != nil {
			return order, errors.New("menu item not found")
		}

		item.MenuItem = menuItem
		item.TotalPrice = float64(item.Quantity) * menuItem.Price
		total += item.TotalPrice

		fullItems = append(fullItems, item)
	}

	order.Total = total
	order.Items = fullItems

	createdOrder, err := uc.orderRepo.CreateOrder(order)
	if err != nil {
		return createdOrder, err
	}

	// Publish order created event
	err = services.PublishOrderCreated(createdOrder)
	if err != nil {
		return createdOrder, err
	}
	return createdOrder, nil
}
