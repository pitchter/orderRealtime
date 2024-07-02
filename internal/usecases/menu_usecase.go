package usecases

import (
	"github.com/pitchter/orderRealtime/internal/adapters/repositories"
	"github.com/pitchter/orderRealtime/internal/entities"
	// "github.com/pitchter/orderRealtime/internal/repositories"
)

type MenuUsecase struct {
    menuRepo repositories.MenuRepository
}

func NewMenuUsecase(repo repositories.MenuRepository) *MenuUsecase {
    return &MenuUsecase{menuRepo: repo}
}

func (uc *MenuUsecase) GetMenu() ([]entities.MenuItem, error) {
    return uc.menuRepo.GetMenu()
}