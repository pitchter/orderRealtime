package entities

import "gorm.io/gorm"

type Order struct {
    gorm.Model
    ID     uint       `gorm:"primaryKey"`
    Items  []MenuItem `gorm:"many2many:order_items;"` // Many-to-many relationship
    Total  float64    `json:"total"`
    Status string     `json:"status"`
}
