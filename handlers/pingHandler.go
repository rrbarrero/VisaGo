package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler endpoint to check service health
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong"})
}
