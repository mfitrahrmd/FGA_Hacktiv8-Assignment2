package models

import "time"

type Order struct {
	OrderID      uint   `gorm:"primaryKey"`
	CustomerName string `binding:"required"`
	OrderedAt    time.Time
	Items        []Item `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
