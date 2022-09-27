package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/config"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/controllers"
	docs "github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

const (
	PORT = 80
)

// @title Orders API
// @version 1.0
// @description This is a simple service for managing orders
// @termsOfService http://swagger.io/terms
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @licence.name Apache 2.0
// @licence.url http://www.apache.org/licence/LICENCE-2.0.html
// @host localhost:80
// @BasePath /
func main() {
	//change argument for below func to your database connection string
	db := config.GetDB(fmt.Sprintf("postgres://developer:developer@localhost:5432/fga_hacktiv8_assignment2?sslmode=disable"))

	//db.Migrate(item.Item{}, order.Order{}) // uncomment this line if run this for the first time to migrate the db

	orderController := controllers.NewOrderController(db.Pg)

	r := gin.Default()

	// localhost/swagger/index.html for swagger API documentation
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// handle request
	r.POST("/orders", orderController.CreateOrder)
	r.GET("/orders", orderController.GetOrders)
	r.PUT("/orders/:id", orderController.UpdateOrder)
	r.DELETE("/orders/:id", orderController.DeleteOrder)

	log.Fatal(r.Run(fmt.Sprintf(":%d", PORT)))
}
