package controllers

import (
	"errors"
	"fmt"
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

func NewOrderController(db *gorm.DB) OrderController {
	return OrderController{
		OrderRepository: order.PostgresOrderRepository{
			DB: db,
		},
		ItemRepository: item.PostgresItemRepository{
			DB: db,
		},
	}
}

// CreateOrder godoc
// @Summary Create new order
// @Description Create new order
// @Tags orders
// @Accept json
// @Produce json
// @Router /orders [post]
func (oc OrderController) CreateOrder(ctx *gin.Context) {
	var reqBody order.Order

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	err = oc.OrderRepository.InsertOne(&reqBody)
	if err != nil {
		helper.ServerError(ctx)

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"orderId": reqBody.OrderID,
		},
	})
}

// GetOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Produce json
// @Router /orders [get]
func (oc OrderController) GetOrders(ctx *gin.Context) {
	orders, err := oc.OrderRepository.FindAll()
	if err != nil {
		helper.ServerError(ctx)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   orders,
	})
}

// UpdateOrder godoc
// @Summary Update existing order
// @Description Update existing order
// @Tags orders
// @Accept json
// @Produce json
// @Router /orders [put]
func (oc OrderController) UpdateOrder(ctx *gin.Context) {
	var reqBody order.Order

	id := ctx.Param("id")
	orderId, _ := strconv.Atoi(id)
	reqBody.OrderID = order.OrderID(orderId)

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	_, err = oc.OrderRepository.FindOne(&order.Order{OrderID: reqBody.OrderID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": fmt.Sprintf("order with id %d was not found", orderId),
			})

			return
		}
		helper.ServerError(ctx)

		return
	}

	for _, itm := range reqBody.Items {
		_, err = oc.ItemRepository.FindOne(&item.Item{ItemID: itm.ItemID})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{
					"status":  "fail",
					"message": fmt.Sprintf("item with id %d was not found", itm.ItemID),
				})

				return
			}
			helper.ServerError(ctx)

			return
		}
	}

	err = oc.OrderRepository.UpdateOne(&reqBody)
	if err != nil {
		helper.ServerError(ctx)

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
// @Router /orders [delete]
func (oc OrderController) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	orderId, _ := strconv.Atoi(id)

	_, err := oc.OrderRepository.FindOne(&order.Order{OrderID: order.OrderID(orderId)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": fmt.Sprintf("order with id %d was not found", orderId),
			})

			return
		}
		helper.ServerError(ctx)

		return
	}

	err = oc.OrderRepository.DeleteOne(&order.Order{
		OrderID: order.OrderID(orderId),
	})
	if err != nil {
		helper.ServerError(ctx)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"orderId": id,
		},
	})

}
