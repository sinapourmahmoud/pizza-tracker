package main

import (
	"log/slog"
	"net/http"
	"pizza-tracker/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Error string
}

type AdminDashboardData struct {
	Orders   []models.Order
	Statuses []string
	Username string
}

func (h *Handler) HandlerLoginGet(c *gin.Context) {

	c.HTML(http.StatusOK, "login.tmpl", LoginData{})

}

func (h *Handler) HandleLoginPost(c *gin.Context) {

	var form struct {
		Username string `form:"username" binding:"required,min=3,max=50"`
		Password string `form:"password" binding:"required,min=8,max=50"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusOK, "login.tmpl", LoginData{Error: "Invalid input" + err.Error()})
		return
	}

	user, err := h.users.AuthenticateUser(form.Username, form.Password)

	if err != nil {
		slog.Info("login failed", "username", form.Username, "error", err)
		c.HTML(http.StatusOK, "login.tmpl", LoginData{
			Error: "Invalid Credentials",
		})

		return
	}
	slog.Info("login successful", "userID", user.ID, "username", user.Username)

	err = SetSessionValue(c, "userID", strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		slog.Error("failed to save userID session", "error", err)
		return
	}

	err = SetSessionValue(c, "username", user.Username)
	if err != nil {
		slog.Error("failed to save username session", "error", err)
		return
	}

	slog.Info("session saved")

	c.Redirect(http.StatusSeeOther, "/admin")

}

func (h *Handler) HandleLogout(c *gin.Context) {

	if err := ClearSession(c); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/login")
}

func (h *Handler) ServeAdminDashboard(c *gin.Context) {
	orders, err := h.orders.GetAllOrders()

	if err != nil {

		c.String(http.StatusInternalServerError, "Error fetching orders")
		return

	}

	username := GetSessionString(c, "username")

	c.HTML(http.StatusOK, "admin.tmpl", AdminDashboardData{
		Orders:   orders,
		Statuses: models.OrderStatuses,
		Username: username,
	})

}

func (h *Handler) HandleOrderPut(c *gin.Context) {

	orderId := c.Param("id")
	newStatus := c.PostForm("status")

	if err := h.orders.UpdateOrderStatus(orderId, newStatus); err != nil {

		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/admin")

}

func (h *Handler) HandleOrderDelete(c *gin.Context) {

	orderID := c.Param("id")

	if err := h.orders.DeleteOrder(orderID); err != nil {

		c.String(http.StatusInternalServerError, err.Error())
		return

	}
	c.Redirect(http.StatusSeeOther, "/admin")

}
