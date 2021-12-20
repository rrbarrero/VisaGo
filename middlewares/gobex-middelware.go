package middlewares

import (
	"rbarrero/visago/gobex"

	"github.com/gin-gonic/gin"
)

func GobexMiddleware(gobex gobex.Gobex) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("gobex", gobex)
		c.Next()
	}
}
