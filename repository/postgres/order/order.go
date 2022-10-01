package order

import (
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/repository/postgres/item"
	"gorm.io/gorm"
	"time"
)

type PostgresOrderRepository struct {
	DB *gorm.DB
}

type OrderID uint

type Order struct {
	OrderID      OrderID     `json:"orderId" gorm:"primaryKey" example:"1"`
	CustomerName string      `json:"customerName" binding:"required" example:"M Fitrah Ramadhan"`
	OrderedAt    time.Time   `json:"orderedAt" binding:"required" example:"2022-09-22T22:00:00+07:00"`
	Items        []item.Item `json:"items" gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CreateOrder struct {
	CustomerName string    `example:"M Fitrah Ramadhan"`
	OrderedAt    time.Time `example:"2022-09-22T22:00:00+07:00"`
	Items        []item.CreateItem
}

type UpdateOrder struct {
	CreateOrder
}

func (por PostgresOrderRepository) InsertOne(order *Order) error {
	if err := por.DB.Create(order).Error; err != nil {
		return err
	}

	return nil
}

func (por PostgresOrderRepository) FindAll() (*[]Order, error) {
	var result []Order

	err := por.DB.Model(Order{}).Preload("Items").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (por PostgresOrderRepository) FindOne(order *Order) (*Order, error) {
	var result Order

	err := por.DB.Model(Order{}).First(&result, order).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (por PostgresOrderRepository) UpdateOne(order *Order) error {
	if err := por.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(order).Error; err != nil {
		return err
	}

	return nil
}

func (por PostgresOrderRepository) DeleteOne(order *Order) error {
	if err := por.DB.Model(&Order{}).Delete(order).Error; err != nil {
		return err
	}

	return nil
}
