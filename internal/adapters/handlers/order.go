package handlers

import (
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
    createdOrder, err := h.orderUsecase.CreateOrder(order)
    if err != nil {
        return utils.HandleError(c, err)
    }
    return c.JSON(createdOrder)
}
