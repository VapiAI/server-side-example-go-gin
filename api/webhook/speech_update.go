package webhook

import (
	"go_gin_app/types"

	"github.com/gin-gonic/gin"
)

func SpeechUpdateHandler(c *gin.Context, payload *types.SpeechUpdatePayload) {
	/**
	 * Handle Business logic here.
	 * Sent during a speech status update during the call. It also lets u know who is speaking.
	 * You can enable this by passing "speech-update" in the serverMessages array while creating the assistant.
	 */

	return
}
