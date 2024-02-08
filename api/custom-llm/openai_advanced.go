package custom_llm

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strings"

	"go_gin_app/config"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

func CustomLLMOpenaiAdvancedHandler(c *gin.Context) {
	var body RequestBody
	envConfig := config.LoadEnvConfig()
	openaiClient := openai.NewClient(envConfig.Openai.ApiKey)
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(500, gin.H{"error": "Invalid request format"})
		return
	}

	ctx := context.Background()

	lastMessage := body.Messages[len(body.Messages)-1]
	promptReq := openai.CompletionRequest{
		Model:       "gpt-3.5-turbo-instruct",
		Prompt:      "Modify given prompt which can act as a prompt template where I put the original prompt and it can modify it according to my intentions so that the final modified prompt is more detailed. You can expand certain terms or keywords.\n----------\nPROMPT: " + lastMessage.Content + ".\nMODIFIED PROMPT: ",
		MaxTokens:   500,
		Temperature: 0.7,
	}
	promptRes, err := openaiClient.CreateCompletion(ctx, promptReq)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	modifiedMessage := append(body.Messages[:len(body.Messages)-1], openai.ChatCompletionMessage{
		Role:    lastMessage.Role,
		Content: strings.Trim(promptRes.Choices[0].Text, "\n"),
	})

	if body.Stream {
		req := openai.ChatCompletionRequest{
			Model:       body.Model,
			Messages:    modifiedMessage,
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

			jsonData, _ := json.Marshal(response.Choices[0].Delta)
			c.Writer.Write([]byte("data: " + string(jsonData) + "\n\n"))
		}
	} else {
		req := openai.ChatCompletionRequest{
			Model:       body.Model,
			Messages:    modifiedMessage,
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
