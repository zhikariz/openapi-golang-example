package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/zhikariz/demo-swagger/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample API for CRUD operations using Swagger.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Load Swagger API documentation

	v1 := r.Group("/api/v1")
	{
		v1.GET("/api/v1/items", GetItems)
		v1.POST("/api/v1/items", CreateItem)
		v1.GET("/api/v1/items/:id", GetItem)
		v1.PUT("/api/v1/items/:id", UpdateItem)
		v1.DELETE("/api/v1/items/:id", DeleteItem)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define your CRUD routes here
	// Example: r.GET("/api/v1/items", getItems)
	//          r.POST("/api/v1/items", createItem)
	//          r.GET("/api/v1/items/:id", getItem)
	//          r.PUT("/api/v1/items/:id", updateItem)
	//          r.DELETE("/api/v1/items/:id", deleteItem)

	r.Run(":8080")
}
