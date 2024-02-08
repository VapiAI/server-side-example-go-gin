package functions

import (
	"github.com/gin-gonic/gin"
)

func BasicHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Basic function API Endpoint - to be implemented",
	})
}
