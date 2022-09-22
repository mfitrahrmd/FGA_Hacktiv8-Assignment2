package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/models"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/services"
	"net/http"
	"strconv"
)

type OrderController struct {
	OrderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return OrderController{
		OrderService: orderService,
	}
}

func (oc OrderController) CreateOrder(ctx *gin.Context) {
	var reqBody models.Order

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})

		return
	}

	err = oc.OrderService.CreateOrder(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "unexpected server error, please try again later",
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("an order has been created with id %d", reqBody.OrderID),
	})
}

func (oc OrderController) GetOrders(ctx *gin.Context) {
	orders, err := oc.OrderService.FindOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": "unexpected server error, please try again later",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   orders,
	})
}

func (oc OrderController) UpdateOrder(ctx *gin.Context) {
	var reqBody models.Order
	orderId := ctx.Param("id")
	atoi, _ := strconv.Atoi(orderId)
	reqBody.OrderID = uint(atoi)

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "please fill all required fields",
		})

		return
	}

	err = oc.OrderService.UpdateOrder(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("updated order with id %d", reqBody.OrderID),
	})
}

func (oc OrderController) DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Param("id")
	atoi, _ := strconv.Atoi(orderId)

	err := oc.OrderService.DeleteOrder(&models.Order{
		OrderID: uint(atoi),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("deleted order with id %s", orderId),
	})

}
