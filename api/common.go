package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Status @Summary Status
// @Description get status
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /status [get]
func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
