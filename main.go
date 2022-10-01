package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/config"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/controllers"
	docs "github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/docs"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/middlewares"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/repository/postgres/item"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-Assignment2/repository/postgres/order"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title Orders API
// @version 1.0
// @description This is a simple service for managing orders
// @termsOfService http://swagger.io/terms
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @licence.name Apache 2.0
// @licence.url http://www.apache.org/licence/LICENCE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	// run 'go run main.go --help' to see required flag for database connection string
	db := config.GetDB(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.PGUSER, config.PGPASSWORD, config.PGHOST, config.PGPORT, config.PGDB))

	db.Migrate(order.Order{}, item.Item{}) // uncomment this line if run this for the first time to migrate the db

	// instantiate order controller
	orderController := controllers.NewOrderController(db.Pg)

	// instantiate gin router
	r := gin.New()

	// localhost/swagger/index.html for swagger API documentation
	// BUG : cannot try Swagger UI that require path parameter, i.e : POST orders that require orderId as path parameter
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// middlewares
	r.Use(middlewares.ErrorHandler())

	// order route handler
	// NOTE : every error that occurs in these controllers will be thrown to error handler middleware, see 'middlewares' folder for the implementation
	rOrder := r.Group("/orders")
	rOrder.POST("", middlewares.ValidatePostOrder, orderController.CreateOrder)
	rOrder.GET("", orderController.GetOrders)
	rOrder.PUT("/:id", middlewares.ValidatePutOrder, orderController.UpdateOrder)
	rOrder.DELETE("/:id", orderController.DeleteOrder)

	// serve the server
	log.Fatal(r.Run(fmt.Sprintf("localhost:%d", config.PORT)))
}
