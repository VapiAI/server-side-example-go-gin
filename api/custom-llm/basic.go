package custom_llm

import (
	"github.com/gin-gonic/gin"
)

func CustomLLMBasicHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Basic custom LLM API Endpoint - to be implemented",
	})
}
