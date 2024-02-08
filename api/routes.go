package api

import (
	custom_llm "go_gin_app/api/custom-llm"
	"go_gin_app/api/functions"
	"go_gin_app/api/webhook"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/api/inbound", InboundHandler)
	r.POST("/api/outbound", OutboundHandler)
	r.POST("/api/webhook", webhook.WebhookHandler)
	r.POST("/api/functions/basic", functions.BasicHandler)
	r.POST("/api/functions/rag", functions.RagHandler)
	r.POST("/api/custom-llm/basic/chat/completions",
		custom_llm.CustomLLMBasicHandler)
	r.POST("/api/custom-llm/openai-sse/chat/completions", custom_llm.CustomLLMOpenaiSSEHandler)
	r.POST("/api/custom-llm/openai-advanced/chat/completions",
		custom_llm.CustomLLMOpenaiAdvancedHandler)
}
