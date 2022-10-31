package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pradeepneosoft/websocket-poc/models"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocketHandler(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	client := models.Client{
		ID:         uuid.Must(uuid.NewRandom()).String(),
		Connection: conn,
	}

	Send(&client, "Server: Welcome! Your ID is "+client.ID)
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			RemoveClient(client)
			return
		}
		ProcessMessage(client, p)
	}

}
func RestHandler(c *gin.Context) {
	var request models.PublishRequest

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println(request)
	Publish(request.Topic, []byte(request.Message))

}
