package controllers

import (
	"Assignment-2/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func NewControllerOrder(db *gorm.DB) *InDB {
	return &InDB{
		DB: db,
	}
}

// GetOrder
// @Summary get all order
// @Description get all order
// @Tags order
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Order
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /orders/{orderId} [get]
func (in *InDB) GetOrder(c *gin.Context) {
	var (
		order  []models.Order
		result gin.H
	)

	err := in.DB.Find(&order).Error
	if err != nil {
		result = gin.H{
			"result": nil,
			"error":  err.Error(),
		}
	}

	if len(order) <= 0 {
		result = gin.H{
			"result": nil,
			"error":  "Data is empty",
		}
	} else {
		result = gin.H{
			"data": order,
		}
	}
	c.JSON(http.StatusOK, result)
}

// CreateOrder
// @Summary create a new order
// @Description create a new order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order body models.Order true "order"
// @Success 200 {object} models.Order
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /orders [post]
func (in *InDB) CreateOrder(c *gin.Context) {
	var order models.Order

	err := json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.DB.Create(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": order,
	})
}

// UpdateOrder
// @Summary update a new order
// @Description update a new order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order body models.Order true "order"
// @Success 200 {object} models.Order
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /orders/{orderId} [put]
func (in *InDB) UpdateOrder(c *gin.Context) {
	var (
		order    models.Order
		newOrder models.Order
	)
	id := c.Param("orderId")

	err := in.DB.First(&order, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.DB.Model(&order).Updates(newOrder).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newOrder,
	})
}

// DeleteOrder
// @Summary delete a new order
// @Description delete a new order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order body models.Order true "order"
// @Success 200 {object} models.Order
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /orders/{orderId} [delete]
func (in *InDB) DeleteOrder(c *gin.Context) {
	var (
		order models.Order
	)
	id := c.Param("orderId")

	err := in.DB.First(&order, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.DB.Delete(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "delete data success !",
	})
}
