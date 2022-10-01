package helper

import (
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/repository/postgres/order"
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type CreatedOrder struct {
	OrderID order.OrderID `json:"orderId" example:"1"`
}

type GetOrders struct {
	Orders []order.Order `json:"orders"`
}

type UpdatedOrder struct {
	OrderID order.OrderID `json:"orderId" example:"1"`
}
type DeletedOrder struct {
	OrderID order.OrderID `json:"orderId" example:"1"`
}

type CreatedOrderResponse struct {
	Status string `json:"status" example:"success"`
	Data   CreatedOrder
}

type GetOrdersResponse struct {
	Status string `json:"status" example:"success"`
	Data   GetOrders
}

type UpdatedOrderResponse struct {
	Status string `json:"status" example:"success"`
	Data   UpdatedOrder
}

type DeletedOrderResponse struct {
	Status string `json:"status" example:"success"`
	Data   DeletedOrder
}
