package main

import (
	"go_gin_app/api"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// envConfig := config.LoadEnvConfig()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	api.SetupRoutes(r)
	r.Run()
}
