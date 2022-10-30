package main

import (
	"github.com/gin-gonic/gin"

	"github.com/pradeepneosoft/websocket-poc/handler"
)

func main() {
	router := gin.Default()
	router.GET("/socket", handler.WebSocketHandler)
	router.POST("/publish", handler.RestHandler)

	router.Run(":8080")
}
