package api

import (
	"github.com/gin-gonic/gin"
)

func OutboundHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Outbound API Endpoint - to be implemented",
	})
}
