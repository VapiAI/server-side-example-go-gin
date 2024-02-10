package webhook

import (
	"encoding/json"
	"go_gin_app/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ...implement GetCallType for other payload types...
func WebhookHandler(c *gin.Context) {
	var genericMessage map[string]interface{}
	if err := c.BindJSON(&genericMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messageData, err := json.Marshal(genericMessage["message"])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var baseMessage types.BaseVapiPayload
	if err := json.Unmarshal(messageData, &baseMessage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	switch baseMessage.Type {
	case types.AssistantRequest:
		var payload types.AssistantRequestPayload
		if err := json.Unmarshal(messageData, &payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error-ass": err.Error()})
			return
		}
		AssistantRequestHandler(c, &payload)

	case types.FunctionCall:
		var payload types.FunctionCallPayload
		if err := json.Unmarshal(messageData, &payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		FunctionCallHandler(c, &payload)

	case types.StatusUpdate:
		var payload types.StatusUpdatePayload
		if err := json.Unmarshal(messageData, &payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		StatusUpdateHandler(c, &payload)

	case types.EndOfCallReport:
		var payload types.EndOfCallReportPayload
		if err := json.Unmarshal(messageData, &payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		EndOfCallReportHandler(c, &payload)

	case types.Hang:
		var payload types.HangPayload
		if err := json.Unmarshal(messageData, &payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		HangEventHandler(c, &payload)

	case types.SpeechUpdate:
		var payload types.SpeechUpdatePayload
		if err := json.Unmarshal(messageData, &payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		SpeechUpdateHandler(c, &payload)
	case types.Transcript:
		var payload types.TranscriptPayload
		if err := json.Unmarshal(messageData, &payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		TranscriptHandler(c, &payload)

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message type"})
	}
}
