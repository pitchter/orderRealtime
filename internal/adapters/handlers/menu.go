package handlers

import (
    "github.com/gofiber/fiber/v2"
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
