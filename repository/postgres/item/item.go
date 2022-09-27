package item

import "gorm.io/gorm"

type PostgresItemRepository struct {
	DB *gorm.DB
}

type Item struct {
	ItemID      uint   `json:"lineItemId" gorm:"primaryKey" example:"1"`
	ItemCode    uint   `json:"itemCode,string" binding:"required" example:"001"`
	Description string `json:"description" binding:"required" example:"Acer Aspire 3"`
	Quantity    int    `json:"quantity" binding:"required" example:"1"`
	OrderID     uint   `json:"orderId" example:"1"`
}

type CreateItem struct {
	ItemCode    uint   `example:"001"`
	Description string `example:"Acer Aspire 3"`
	Quantity    int    `example:"1"`
}

func (por PostgresItemRepository) FindOne(item *Item) (*Item, error) {
	var result Item

	err := por.DB.Model(Item{}).First(&result, item).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
