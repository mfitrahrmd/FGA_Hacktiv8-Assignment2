package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/helper"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/repository/postgres/order"
)

func ValidatePostOrder(ctx *gin.Context) {
	var reqBody order.Order

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.Error(helper.INVALID_REQUEST_BODY)
		ctx.Abort()

		return
	}

	ctx.Set("reqBody", reqBody)
}

func ValidatePutOrder(ctx *gin.Context) {
	ValidatePostOrder(ctx)
}
