package custom_llm

import (
	"github.com/gin-gonic/gin"
)

func CustomLLMOpenaiSSEHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OpenAI SSE custom LLM API Endpoint - to be implemented",
	})
}
