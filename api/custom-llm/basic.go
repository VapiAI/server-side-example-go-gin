package custom_llm

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ChatCompletionRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages,omitempty"`
	MaxTokens   *int      `json:"max_tokens,omitempty"`
	Temperature *float32  `json:"temperature,omitempty"`
	Stream      *bool     `json:"stream,omitempty"`
	Call        *bool     `json:"call,omitempty"`
}

type Message struct {
	Content *string `json:"content,omitempty"`
}

type ChatCompletionResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	SystemFingerprint *string  `json:"system_fingerprint,omitempty"`
	Choices           []Choice `json:"choices"`
}

type Choice struct {
	Index        int    `json:"index"`
	Delta        Delta  `json:"delta"`
	Logprobs     *int   `json:"logprobs,omitempty"`
	FinishReason string `json:"finish_reason"`
}

type Delta struct {
	Content string `json:"content"`
}

func CustomLLMBasicHandler(c *gin.Context) {
	response := ChatCompletionResponse{
		ID:      "chatcmpl-8mcLf78g0quztp4BMtwd3hEj58Uof",
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   "gpt-3.5-turbo-0613",
		Choices: []Choice{
			{
				Index: 0,
				Delta: Delta{
					Content: "I am a highly intelligent question-answering AI. I can help you with any question you have.",
				},
				FinishReason: "stop",
			},
		},
	}

	c.JSON(201, response)
}
