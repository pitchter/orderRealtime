package repositories

import (
    "context"
    "encoding/json"
    "github.com/pitchter/orderRealtime/internal/entities"
    "github.com/pitchter/orderRealtime/internal/adapters/redis"
    "gorm.io/gorm"
)

type MenuRepository interface {
    GetMenu() ([]entities.MenuItem, error)
    GetMenuItemByID(id uint) (entities.MenuItem, error)
    CreateMenuItem(item entities.MenuItem) (entities.MenuItem, error)
}

type menuRepository struct {
    db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
    return &menuRepository{db: db}
}

// ฟังก์ชั่นดึงข้อมูลเมนู
func (repo *menuRepository) GetMenu() ([]entities.MenuItem, error) {
    // ลองดึงข้อมูลจาก Redis ก่อน
    val, err := redis.Rdb.Get(context.Background(), "menu").Result()
    if err == redis.Nil {
        // ถ้าไม่มีข้อมูลใน Redis (redis.Nil) ให้ดึงจากฐานข้อมูล
        var menu []entities.MenuItem
        result := repo.db.Find(&menu)
        if result.Error != nil {
            return nil, result.Error
        }
        // เก็บข้อมูลใน Redis เพื่อใช้ในครั้งต่อไป
        jsonMenu, _ := json.Marshal(menu)
        redis.Rdb.Set(context.Background(), "menu", jsonMenu, 0)
        return menu, nil
    } else if err != nil {
        return nil, err
    }

    // ถ้ามีข้อมูลใน Redis ให้แปลงจาก JSON กลับมาเป็นโครงสร้างข้อมูลเดิม
    var menu []entities.MenuItem
    json.Unmarshal([]byte(val), &menu)
    return menu, nil
}

func (repo *menuRepository) GetMenuItemByID(id uint) (entities.MenuItem, error) {
    var menuItem entities.MenuItem
    result := repo.db.First(&menuItem, id)
    return menuItem, result.Error
}

// ฟังก์ชั่นสร้างเมนูใหม่
func (repo *menuRepository) CreateMenuItem(item entities.MenuItem) (entities.MenuItem, error) {
    result := repo.db.Create(&item)
    if result.Error != nil {
        return item, result.Error
    }

    // ลบข้อมูล cache ใน Redis เพื่อให้ข้อมูลอัปเดตในครั้งต่อไป
    redis.Rdb.Del(context.Background(), "menu")
    return item, nil
}
