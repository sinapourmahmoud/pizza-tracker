package main

import (
	"net/http"
	"pizza-tracker/internal/models"

	"github.com/gin-gonic/gin"
)

type OrderFormData struct {
	PizzaTypes []string
	PizzaSizes []string
}

type OrderRequest struct {
	Name    string `form:"name" binding:"required,min=2,max=100"`
	Phone   string `form:"phone" binding:"required,min=10,max=20"`
	Address string `form:"address" binding:"required,min=5,max=200"`
	PizzaSizes   []string `form:"size" binding:"required,min=1,dive,valid_pizza_size"`
	PizzaTypes   []string `form:"pizza" binding:"required,min=1,dive,valid_pizza_type"`
	Instructions []string `form:"instructions" binding:"max=200"`
}


func (h *Handler) ServeNewOrderPost(c *gin.Context){
	c.HTML(http.StatusOK,"order.tmpl",OrderFormData{
		PizzaTypes: models.PizzaTypes,
		PizzaSizes: models.PizzaSizes,
	})
}