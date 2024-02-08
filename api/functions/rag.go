package functions

import (
	"github.com/gin-gonic/gin"
)

func RagHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "RAG function API Endpoint - to be implemented",
	})
}
