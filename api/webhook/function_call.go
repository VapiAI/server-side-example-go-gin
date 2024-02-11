package webhook

import (
	"fmt"
	"go_gin_app/functions"
	"go_gin_app/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type FunctionCallPayload struct {
	FunctionCall struct {
		Name       string                 `json:"name"`
		Parameters map[string]interface{} `json:"parameters"`
	} `json:"functionCall"`
}

func FunctionCallHandler(c *gin.Context, payload *types.FunctionCallPayload) {

	fmt.Printf("Payload: %+v\n", payload)

	if payload.FunctionCall.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request."})
		return
	}

	functionName := payload.FunctionCall.Name
	parameters := payload.FunctionCall.Parameters

	switch functionName {
	case "getRandomName":
		var params functions.NameParams
		mapstructure.Decode(parameters, &params)
		result, err := functions.GetRandomName(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"result": result})

	case "getCharacterInspiration":
		var params functions.GetCharacterInspirationParams
		mapstructure.Decode(parameters, &params)
		result := functions.GetCharacterInspiration(params)
		c.JSON(http.StatusCreated, gin.H{"result": result.Result, "forwardToClientEnabled": result.ForwardToClientEnabled})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid function name."})
		return
	}

}
