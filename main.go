package main

import (
	"github.com/flame/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	handler.RouterHandler(server)

	server.Run(":8000")
}
