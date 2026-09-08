package main

import (
	"log/slog"
	"net/http"
	"pizza-tracker/internal/models"

	"github.com/gin-gonic/gin"
)

type OrderFormData struct {
	PizzaTypes []string
	PizzaSizes []string
}

type OrderRequest struct {
	Name         string   `form:"name" binding:"required,min=2,max=100"`
	Phone        string   `form:"phone" binding:"required,min=10,max=20"`
	Address      string   `form:"address" binding:"required,min=5,max=200"`
	PizzaSizes   []string `form:"size" binding:"required,min=1,dive,valid_pizza_size"`
	PizzaTypes   []string `form:"pizza" binding:"required,min=1,dive,valid_pizza_type"`
	Instructions []string `form:"instructions" binding:"max=200"`
}

func (h *Handler) ServeNewOrderForm(c *gin.Context) {
	c.HTML(http.StatusOK, "order.tmpl", OrderFormData{
		PizzaTypes: models.PizzaTypes,
		PizzaSizes: models.PizzaSizes,
	})
}

func (h *Handler) HandleNewOrderPost(c *gin.Context) {

	var form OrderRequest

	if err := c.ShouldBind(form); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItems := make([]models.OrderItem, len(form.PizzaSizes))

	for i := range orderItems {
		orderItems[i] = models.OrderItem{
			Size:         form.PizzaSizes[i],
			Pizza:        form.PizzaTypes[i],
			Instructions: form.Instructions[i],
		}
	}

	order := models.Order{
		CustomerName: form.Name,
		Phone:        form.Phone,
		Address:      form.Address,
		Status:       models.OrderStatuses[0],
		Items:        orderItems,
	}

	if err := h.orders.CreateOrder(&order); err != nil {
		slog.Error("Failed to create order", "error", err)
		c.String(http.StatusInternalServerError, "Something went wrong")
		return
	}

	slog.Info("Order created", "orderId", order.ID, "customer", order.CustomerName)

	c.Redirect(http.StatusSeeOther, "/customer/"+order.ID)

}

func (h *Handler) serveCustomer(c *gin.Context) {
	orderId := c.Param("id")

	if orderId == "" {
		c.String(http.StatusBadRequest, "Order ID is required")

	}

	order, err := h.orders.GetOrder(orderId)

	if err != nil {
		c.String(http.StatusNotFound, "Order not found")
		return
	}
	c.HTML(http.StatusOK, "customer.tmpl", gin.H{
		"Order": order,
	})

}
