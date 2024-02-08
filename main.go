package main

import (
	"go_gin_app/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.SetupRoutes(r)
	r.Run()
}
