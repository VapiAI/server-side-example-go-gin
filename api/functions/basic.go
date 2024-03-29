package functions

import (
	"encoding/json"
	"go_gin_app/functions"
	"go_gin_app/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func BasicHandler(c *gin.Context) {
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
	case types.FunctionCall:
		var payload types.FunctionCallPayload
		if err := json.Unmarshal(messageData, &payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if payload.FunctionCall.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request."})
			return
		}

		functionName := payload.FunctionCall.Name
		parameters := payload.FunctionCall.Parameters

		if functionName == "getRandomName" {

			var params functions.NameParams
			mapstructure.Decode(parameters, &params)
			result, err := functions.GetRandomName(params)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, gin.H{"result": result})
		}
	default:
	}
}
