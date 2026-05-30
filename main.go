package main

import (
	"github.com/gin-gonic/gin"
	"russ-bassett-orders-api/routes"
)

func main() {

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}