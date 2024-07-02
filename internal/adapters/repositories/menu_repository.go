package repositories

import (
    "context"
    "encoding/json"
    "github.com/pitchter/orderRealtime/internal/entities"
    "github.com/pitchter/orderRealtime/internal/adapters/redis"
 
)

type MenuRepository interface {
    GetMenu() ([]entities.MenuItem, error)
}

type menuRepository struct{}

func NewMenuRepository() MenuRepository {
    return &menuRepository{}
}

func (repo *menuRepository) GetMenu() ([]entities.MenuItem, error) {
    val, err := redis.Rdb.Get(context.Background(), "menu").Result()
    if err == nil {
        menu := []entities.MenuItem{
            {ID: 1, Name: "Pizza", Price: 10.99, Category: "Food"},
            {ID: 2, Name: "Burger", Price: 8.99, Category: "Food"},
            {ID: 3, Name: "Coke", Price: 1.99, Category: "Drink"},
        }
        jsonMenu, _ := json.Marshal(menu)
        redis.Rdb.Set(context.Background(), "menu", jsonMenu, 0)
        return menu, nil
    } else if err != nil {
        return nil, err
    }

    var menu []entities.MenuItem
    json.Unmarshal([]byte(val), &menu)
    return menu, nil
}
