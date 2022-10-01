package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/helper"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/repository"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/repository/postgres/item"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/repository/postgres/order"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type OrderController struct {
	OrderRepository repository.OrderRepository
	ItemRepository  repository.ItemRepository
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{
		OrderRepository: order.PostgresOrderRepository{
			DB: db,
		},
		ItemRepository: item.PostgresItemRepository{
			DB: db,
		},
	}
}

//

// CreateOrder godoc
// @Summary Create an order
// @Description Create an order, and return the order id
// @Tags orders
// @Accept json
// @Produce json
// @Param body body order.CreateOrder true "order data to create"
// @Success 200 {object} helper.CreatedOrderResponse
// @Failure 500 {object} helper.ServerErrorResponse
// @Router /orders [post]
func (oc OrderController) CreateOrder(ctx *gin.Context) {
	body, ok := ctx.Get("reqBody")
	if !ok {
		ctx.Error(errors.New("cannot get request body from context"))

		return
	}

	reqBody := body.(order.Order)

	err := oc.OrderRepository.InsertOne(&reqBody)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusCreated, helper.CreatedOrderResponse{
		Status: "success",
		Data: helper.CreatedOrder{
			OrderID: reqBody.OrderID,
		},
	})
}

// GetOrders godoc
// @Summary Get all orders
// @Description Get details of all orders
// @Tags orders
// @Produce json
// @Success 200 {object} helper.GetOrdersResponse
// @Failure 500 {object} helper.ServerErrorResponse
// @Router /orders [get]
func (oc OrderController) GetOrders(ctx *gin.Context) {
	orders, err := oc.OrderRepository.FindAll()
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, helper.GetOrdersResponse{
		Status: "success",
		Data: helper.GetOrders{
			Orders: *orders,
		},
	})
}

// UpdateOrder godoc
// @Summary Update an order
// @Description Update existing order with new data
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param body body order.UpdateOrder true "order data to update"
// @Success 200 {object} helper.UpdatedOrderResponse
// @Failure 404 {object} helper.NotFoundResponse
// @Failure 400 {object} helper.BadRequestResponse
// @Failure 500 {object} helper.ServerErrorResponse
// @Router /orders [put]
func (oc OrderController) UpdateOrder(ctx *gin.Context) {
	body, ok := ctx.Get("reqBody")
	if !ok {
		ctx.Error(errors.New("cannot get request body from context"))

		return
	}

	reqBody := body.(order.Order)

	id := ctx.Param("id")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(err)

		return
	}

	reqBody.OrderID = order.OrderID(orderId)

	_, err = oc.OrderRepository.FindOne(&order.Order{OrderID: reqBody.OrderID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Error(helper.NOT_FOUND)

			return
		}
		ctx.Error(err)

		return
	}

	for _, itm := range reqBody.Items {
		_, err = oc.ItemRepository.FindOne(&item.Item{ItemID: itm.ItemID})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.Error(helper.NOT_FOUND)

				return
			}
			ctx.Error(err)

			return
		}
	}

	err = oc.OrderRepository.UpdateOne(&reqBody)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"orderId": reqBody.OrderID,
		},
	})
}

// DeleteOrder godoc
// @Summary Delete existing order
// @Description Delete existing order
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} helper.DeletedOrderResponse
// @Failure 404 {object} helper.NotFoundResponse
// @Failure 500 {object} helper.ServerErrorResponse
// @Router /orders [delete]
func (oc OrderController) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(err)

		return
	}

	_, err = oc.OrderRepository.FindOne(&order.Order{OrderID: order.OrderID(orderId)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Error(helper.NOT_FOUND)

			return
		}
		ctx.Error(err)

		return
	}

	err = oc.OrderRepository.DeleteOne(&order.Order{
		OrderID: order.OrderID(orderId),
	})
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"orderId": id,
		},
	})

}
