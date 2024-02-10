package webhook

import (
	"go_gin_app/types"

	"github.com/gin-gonic/gin"
)

func TranscriptHandler(c *gin.Context, payload *types.TranscriptPayload) {
	// Handle Business logic here.
	// Sent during a call whenever the transcript is available for certain chunk in the stream.
	// You can store the transcript in your database or have some other business logic.

}
