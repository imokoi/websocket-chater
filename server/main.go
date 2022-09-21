package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/olahol/melody"
)

var m *melody.Melody

func main() {
	r := gin.Default()
	m = melody.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ws", func(c *gin.Context) {
		_ = m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(connectionHandler)
	m.HandleDisconnect(disconnectionHandler)
	_ = r.Run(":8888")
}

// connectionHandler is called when a new websocket connection is established.
func connectionHandler(s *melody.Session) {
	id := uuid.NewString()
	s.Set("id", id)
	_ = s.Write([]byte("Hello"))
}

func disconnectionHandler(s *melody.Session) {
	idObject, ok := s.Get("id")
	if !ok {
		return
	}
	id := idObject.(string)
	fmt.Println(id)
	_ = s.Write([]byte("Bye"))
	_ = m.BroadcastOthers([]byte("Bye"), s)
}
