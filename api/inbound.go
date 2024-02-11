package api

import (
	"go_gin_app/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InboundHandler(c *gin.Context) {
	var reqBody struct {
		Message struct {
			Type types.VapiWebhookEnum `json:"type"`
		} `json:"message"`
	}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch reqBody.Message.Type {
	case types.AssistantRequest:
		name := "Paula"
		modelName := "gpt-3.5-turbo"
		temp := 0.7
		systemPrompt := "You're Paula, an AI assistant who can help user draft beautiful emails to their clients based on the user requirements. Then Call sendEmail function to actually send the email."
		functionDescription := "Send email to the given email address and with the given content."
		firstMessage := "Hi, I'm Paula, your personal email assistant."

		assistant := &types.Assistant{
			Name: &name,
			Model: &types.Model{
				Provider:     "openai",
				Model:        modelName,
				Temperature:  &temp,
				SystemPrompt: &systemPrompt,
				Functions: []types.Function{
					{
						Name:        "sendEmail",
						Description: &functionDescription,
						Parameters: map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"email": map[string]interface{}{
									"type":        "string",
									"description": "Email to which we want to send the content.",
								},
								"content": map[string]interface{}{
									"type":        "string",
									"description": "Actual Content of the email to be sent.",
								},
							},
							"required": []string{"email"},
						},
					},
				},
			},
			Voice: &types.Voice{
				Provider: "11labs",
				VoiceId:  "paula",
			},
			FirstMessage: &firstMessage,
		}
		c.JSON(http.StatusCreated, gin.H{"assistant": assistant})
	default:
		// Ignoring invalid message type without sending an error
		c.JSON(http.StatusCreated, gin.H{})
	}
}
