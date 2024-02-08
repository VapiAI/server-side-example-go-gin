package custom_llm

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"go_gin_app/config"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type RequestBody struct {
	Model       string                         `json:"model"`
	Messages    []openai.ChatCompletionMessage `json:"messages"`
	MaxTokens   int                            `json:"max_tokens"`
	Temperature float32                        `json:"temperature"`
	Stream      bool                           `json:"stream"`
}

var envConfig = config.LoadEnvConfig()
var openaiClient = openai.NewClient(envConfig.Openai.ApiKey)

func CustomLLMOpenaiSSEHandler(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(500, gin.H{"error": "Invalid request format"})
		return
	}

	ctx := context.Background()

	if body.Stream {
		req := openai.ChatCompletionRequest{
			Model:       body.Model,
			Messages:    body.Messages,
			MaxTokens:   body.MaxTokens,
			Temperature: body.Temperature,
			Stream:      true,
		}
		completionStream, err := openaiClient.CreateChatCompletionStream(ctx, req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer completionStream.Close()

		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		for {
			response, err := completionStream.Recv()
			if errors.Is(err, io.EOF) {
				c.Writer.Flush()
				return
			}

			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			jsonData, _ := json.Marshal(response)
			c.Writer.Write([]byte("data: " + string(jsonData) + "\n\n"))
		}
	} else {
		// Non-streaming completion logic here
		req := openai.ChatCompletionRequest{
			Model:       body.Model,
			Messages:    body.Messages,
			MaxTokens:   body.MaxTokens,
			Temperature: body.Temperature,
			Stream:      false,
		}
		response, err := openaiClient.CreateChatCompletion(ctx, req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, response)
	}
}
