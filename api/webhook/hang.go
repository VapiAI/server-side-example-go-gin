package webhook

import (
	"go_gin_app/types"

	"github.com/gin-gonic/gin"
)

func HangEventHandler(c *gin.Context, payload *types.HangPayload) {
	/**
	 * Handle Business logic here.
	 * Sent once the call is terminated by user.
	 * You can update the database or have some followup actions or workflow triggered.
	 */
	return
}
