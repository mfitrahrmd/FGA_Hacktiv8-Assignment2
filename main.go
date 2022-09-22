package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/config"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/controllers"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/models"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/services"
	"log"
)

func main() {
	//change argument for below func to your database connection string
	db := config.GetDB(fmt.Sprintf("postgres://user:password@localhost:5432/database?sslmode=disable"))

	err := db.AutoMigrate(models.Order{}, models.Item{})
	if err != nil {
		log.Fatal(err.Error())
	}

	orderService := services.NewOrderService(db)
	orderController := controllers.NewOrderController(orderService)

	r := gin.Default()

	r.POST("/orders", orderController.CreateOrder)
	r.GET("/orders", orderController.GetOrders)
	r.PUT("/orders/:id", orderController.UpdateOrder)
	r.DELETE("/orders/:id", orderController.DeleteOrder)

	log.Fatal(r.Run(":80"))

}
