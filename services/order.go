package services

import (
	"errors"
	"fmt"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/models"
	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(database *gorm.DB) OrderService {
	return OrderService{
		db: database,
	}
}

func (oS OrderService) CreateOrder(order *models.Order) error {
	if err := oS.db.Create(order).Error; err != nil {
		return err
	}

	return nil
}

func (oS OrderService) FindOrders() (*[]models.Order, error) {
	var result []models.Order

	err := oS.db.Model(models.Order{}).Preload("Items").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (oS OrderService) isOrderExist(order *models.Order) bool {
	var result []models.Order

	oS.db.Model(models.Order{}).Where("order_id = ?", order.OrderID).Find(&result)

	if len(result) < 1 {
		fmt.Println(false)
		return false
	}

	return true
}

func (oS OrderService) isItemExist(item models.Item) bool {
	var result []models.Item

	oS.db.Model(models.Item{}).Where("item_id = ?", item.ItemID).Find(&result)

	if len(result) < 1 {
		fmt.Println(false)
		return false
	}

	return true
}

func (oS OrderService) UpdateOrder(order *models.Order) error {
	if !oS.isOrderExist(order) {
		return errors.New(fmt.Sprintf("order with id %d was not found", order.OrderID))
	}

	for _, item := range order.Items {
		if !oS.isItemExist(item) {
			return errors.New(fmt.Sprintf("item with id %d was not found", item.ItemID))
		}
	}

	if err := oS.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(order).Error; err != nil {
		panic(err)
	}

	return nil
}

func (oS OrderService) DeleteOrder(order *models.Order) error {
	if !oS.isOrderExist(order) {
		return errors.New(fmt.Sprintf("order with id %d was not found", order.OrderID))
	}

	result := oS.db.Delete(order)
	if err := result.Error; err != nil {
		panic(err)
	}

	return nil
}
