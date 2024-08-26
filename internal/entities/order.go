package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID           uint       `gorm:"primaryKey"`
	CustomerName string     `json:"customer_name"`
	TableNumber  int        `json:"table_number"`
	Items        []OrderItem `gorm:"foreignKey:OrderID"`
	Total        float64    `json:"total"`
	Status       string     `json:"status"`
}

type OrderItem struct {
    gorm.Model
    OrderID     uint      `json:"order_id"`
    MenuItemID  uint      `json:"menu_item_id"`
    Quantity    int       `json:"quantity"`
    MenuItem    MenuItem  `json:"menu_item"` // อ้างอิง MenuItem เพื่อดึงข้อมูล
    TotalPrice  float64   `json:"total_price"`
}
