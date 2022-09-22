package models

type Item struct {
	ItemID      uint   `json:"lineItemId" gorm:"primaryKey"`
	ItemCode    uint   `json:"itemCode,string" binding:"required"`
	Description string `binding:"required"`
	Quantity    int    `binding:"required"`
	OrderID     uint
}
