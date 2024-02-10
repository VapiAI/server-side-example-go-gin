package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_gin_app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	PhoneNumberId  string `json:"phoneNumberId"`
	AssistantId    string `json:"assistantId"`
	CustomerNumber string `json:"customerNumber"`
}

func OutboundHandler(c *gin.Context) {
	envConfig := config.LoadEnvConfig()
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf(":Request Body: \n")

	requestBodyBytes, _ := json.Marshal(requestBody)

	fmt.Println(requestBody)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/call/phone", envConfig.Vapi.BaseUrl), bytes.NewBuffer(requestBodyBytes))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", envConfig.Vapi.ApiKey))

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	fmt.Println("requestBody end", resp.StatusCode, http.StatusOK)

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("HTTP error! status: %d", resp.StatusCode)})
		return
	}

	var data interface{}
	json.NewDecoder(resp.Body).Decode(&data)

	c.JSON(http.StatusOK, data)
}
