package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"russ-bassett-orders-api/models"
	"russ-bassett-orders-api/storage"
)

func CreateOrder(c *gin.Context) {

	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if order.ClientName == "" || order.ProjectType == "" || order.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Required fields are missing",
		})
		return
	}

	validStatus := map[string]bool{
		"Pending":     true,
		"In Progress": true,
		"Completed":   true,
	}

	if !validStatus[order.Status] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Status must be Pending, In Progress, or Completed",
		})
		return
	}

	order.ID = len(storage.Orders) + 1

	storage.Orders = append(storage.Orders, order)

	c.JSON(http.StatusCreated, order)
}

func GetOrder(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	for _, order := range storage.Orders {

		if order.ID == id {
			c.JSON(http.StatusOK, order)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Order not found",
	})
}

func ListOrders(c *gin.Context) {

	status := c.Query("status")

	var result []models.Order

	if status == "" {
		c.JSON(http.StatusOK, storage.Orders)
		return
	}

	for _, order := range storage.Orders {

		if order.Status == status {
			result = append(result, order)
		}
	}

	c.JSON(http.StatusOK, result)
}

func UpdateOrder(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var updatedOrder models.Order

	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	validStatus := map[string]bool{
		"Pending":     true,
		"In Progress": true,
		"Completed":   true,
	}

	if !validStatus[updatedOrder.Status] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Status must be Pending, In Progress, or Completed",
		})
		return
	}

	for index, order := range storage.Orders {

		if order.ID == id {

			updatedOrder.ID = id

			storage.Orders[index] = updatedOrder

			c.JSON(http.StatusOK, updatedOrder)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Order not found",
	})
}

func DeleteOrder(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	for index, order := range storage.Orders {

		if order.ID == id {

			storage.Orders = append(
				storage.Orders[:index],
				storage.Orders[index+1:]...,
			)

			c.JSON(http.StatusOK, gin.H{
				"message": "Order deleted successfully",
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Order not found",
	})
}