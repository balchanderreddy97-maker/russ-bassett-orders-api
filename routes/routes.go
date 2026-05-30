package routes

import (
	"github.com/gin-gonic/gin"
	"russ-bassett-orders-api/handlers"
	"russ-bassett-orders-api/middleware"
)

func SetupRoutes(router *gin.Engine) {

	router.Use(middleware.APIKeyAuth())

	router.POST("/orders", handlers.CreateOrder)
	router.GET("/orders", handlers.ListOrders)
	router.GET("/orders/:id", handlers.GetOrder)
	router.PUT("/orders/:id", handlers.UpdateOrder)
	router.DELETE("/orders/:id", handlers.DeleteOrder)
}