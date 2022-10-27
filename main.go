package main

import (
	"Assignment-2/config"
	"Assignment-2/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

const (
	_port_ = ":8080"
)

// @title Assigment-2
// @version 1.0
// @description This is a sample server Petstore server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9000
// @BasePath /
func main() {
	db := config.StartDB()
	if db != nil {
		fmt.Println("Running database is failed")
	}

	orderController := controllers.NewControllerOrder(db)
	itemController := controllers.NewControllerItem(db)

	router := gin.Default()

	router.GET("/orders", orderController.GetOrder)
	router.POST("/orders", orderController.CreateOrder)
	router.PUT("/orders/:orderId", orderController.UpdateOrder)
	router.DELETE("/orders/:orderId", orderController.DeleteOrder)

	router.GET("/items", itemController.GetItem)
	router.POST("/items", itemController.CreateItem)
	router.PUT("/items/:itemId", itemController.UpdateItem)
	router.DELETE("/items/:itemId", itemController.DeleteItem)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(_port_)
}
