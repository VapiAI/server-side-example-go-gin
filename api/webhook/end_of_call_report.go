package webhook

import (
	"go_gin_app/types"

	"github.com/gin-gonic/gin"
)

func EndOfCallReportHandler(c *gin.Context, payload *types.EndOfCallReportPayload) {
	// Handle business logic here
	// You can store the information like summary, recordingUrl, or even the full messages list in the database

	return
}
