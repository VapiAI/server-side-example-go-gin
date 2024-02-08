package webhook

import (
	"github.com/gin-gonic/gin"
)

func WebhookHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Webhook API Endpoint - to be implemented",
	})
}
