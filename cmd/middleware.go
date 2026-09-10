package main

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)
func (h *Handler) AuthMiddleware() gin.HandlerFunc {

    return func(c *gin.Context) {

        userID := GetSessionString(c, "userID")

        slog.Info("auth middleware",
            "userID", userID,
        )

        if userID == "" {
            slog.Info("no userID in session")

            c.Redirect(http.StatusSeeOther, "/login")
            c.Abort()
            return
        }

        _, err := h.users.GetUserByID(userID)

        if err != nil {
            slog.Error("user not found", "userID", userID, "error", err)

            ClearSession(c)
            c.Redirect(http.StatusSeeOther, "/login")
            c.Abort()
            return
        }

        c.Next()
    }
}