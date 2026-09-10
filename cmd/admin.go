package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Error string
}

type AdminDashboardData struct {
	Username string
}

func (h *Handler) HandlerLoginGet(c *gin.Context) {

	c.HTML(http.StatusOK, "login.tmpl", LoginData{})

}

func (h *Handler) HandleLoginPost(c *gin.Context) {

	var form struct {
		Username string `form:"username" binding:"required,min=3,max=50"`
		Password string `form:"username" binding:"required,min,max="`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusOK, "login.tmpl", LoginData{Error: "Invalid input" + err.Error()})
		return
	}

	user, err := h.users.AuthenticateUser(form.Username, form.Password)

	if err != nil {
		c.HTML(http.StatusOK, "login.tmpl", LoginData{
			Error: "Invalid Credentials",
		})

		return
	}

	SetSessionValue(c, "userID", user.ID)
	SetSessionValue(c, "username", user.Username)
	c.Redirect(http.StatusSeeOther, "/admin")

}

func (h *Handler) HandleLogout(c *gin.Context) {

	if err := ClearSession(c); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/login")
}


func (h *Handler) ServeAdminDashboard(c *gin.Context){

	username := GetSessionString(c,"username")

	c.HTML(http.StatusOK,"admin.tmpl",AdminDashboardData{
		Username: username,
	})


}


