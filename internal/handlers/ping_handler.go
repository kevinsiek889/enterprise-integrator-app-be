package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler godoc
// @Summary Ping test
// @Description Returns pong
// @Tags Ping
// @Success 200 {string} string "pong"
// @Router /ping [get]
func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}