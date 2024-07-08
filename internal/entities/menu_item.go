package entities

import "gorm.io/gorm"

type MenuItem struct {
    gorm.Model
    ID       uint    `gorm:"primaryKey"`
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Category string  `json:"category"`
}
