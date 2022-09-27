package repository

import (
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/repository/postgres/item"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/repository/postgres/order"
)

type OrderRepository interface {
	InsertOne(*order.Order) error
	FindAll() (*[]order.Order, error)
	FindOne(*order.Order) (*order.Order, error)
	UpdateOne(*order.Order) error
	DeleteOne(*order.Order) error
}

type ItemRepository interface {
	FindOne(*item.Item) (*item.Item, error)
}
