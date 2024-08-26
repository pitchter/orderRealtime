package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/pitchter/orderRealtime/internal/entities"
	"github.com/pitchter/orderRealtime/internal/usecases"
	"github.com/pitchter/orderRealtime/internal/utils"
)

type OrderHandler struct {
	orderUsecase *usecases.OrderUsecase
}

func NewOrderHandler(orderUsecase *usecases.OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUsecase: orderUsecase}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order entities.Order
	if err := c.BodyParser(&order); err != nil {
		return utils.HandleError(c, err)
	}

	if order.CustomerName == "" {
		return utils.HandleError(c, errors.New("customer name is required"))
	}

	 // ดึงเลขโต๊ะจาก query parameter
	//  tableNumber, err := c.QueryInt("table", 0)
	//  if err != nil || tableNumber <= 0 {
	// 	 return utils.HandleError(c, errors.New("invalid table number"))
	//  }
 
	 // ตั้งค่า table number ในคำสั่งซื้อ
	//  order.TableNumber = tableNumber

	if order.TableNumber <= 0 {
        return utils.HandleError(c, errors.New("invalid table number"))
    }


    for _, item := range order.Items {
        if item.Quantity <= 0 {
            return utils.HandleError(c, errors.New("quantity must be greater than 0"))
        }
    }

	createdOrder, err := h.orderUsecase.CreateOrder(order)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.JSON(createdOrder)
}
