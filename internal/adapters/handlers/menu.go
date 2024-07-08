package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/pitchter/orderRealtime/internal/entities"
    "github.com/pitchter/orderRealtime/internal/usecases"
    "github.com/pitchter/orderRealtime/internal/utils"
)

type MenuHandler struct {
    menuUsecase *usecases.MenuUsecase
}

func NewMenuHandler(menuUsecase *usecases.MenuUsecase) *MenuHandler {
    return &MenuHandler{menuUsecase: menuUsecase}
}

func (h *MenuHandler) GetMenu(c *fiber.Ctx) error {
    menu, err := h.menuUsecase.GetMenu()
    if err != nil {
        return utils.HandleError(c, err)
    }
    return c.JSON(menu)
}

func (h *MenuHandler) CreateMenuItem(c *fiber.Ctx) error {
    var item entities.MenuItem
    if err := c.BodyParser(&item); err != nil {
        return utils.HandleError(c, err)
    }
    createdItem, err := h.menuUsecase.CreateMenuItem(item)
    if err != nil {
        return utils.HandleError(c, err)
    }
    return c.JSON(createdItem)
}
