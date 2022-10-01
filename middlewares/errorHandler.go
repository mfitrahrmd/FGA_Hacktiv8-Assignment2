package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/helper"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		for _, e := range ctx.Errors {
			err := e.Err

			if myErr, ok := err.(helper.MyError); ok {
				ctx.AbortWithStatusJSON(myErr.Code, gin.H{
					"status":  "fail",
					"message": myErr.Msg,
					"data":    myErr.Data,
				})
			} else {
				fmt.Println(e.Error())

				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"status":  "fail",
					"message": "Unexpected server error",
				})
			}
		}
	}
}
