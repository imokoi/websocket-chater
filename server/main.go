package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/olahol/melody"
)

var m *melody.Melody
var sessionMap sync.Map

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
	m.HandleMessage(messageHandler)
	m.HandleDisconnect(disconnectionHandler)
	m.HandlePong(pongHandler)
	_ = r.Run(":8888")
}

// pongHandler is called when a pong message is received from a client.
func pongHandler(s *melody.Session) {
	fmt.Println("pong")
	_ = s.Write([]byte("Pong"))
}

// connectionHandler is called when a new websocket connection is established.
func connectionHandler(s *melody.Session) {
	id := uuid.NewString()
	sessionMap.Store(id, s)
	s.Set("id", id)
	// replay the message to the client
	_ = s.Write([]byte("Hello"))
}

// messageHandler is called when a message is received from a client.
func messageHandler(s *melody.Session, msg []byte) {
	fmt.Println(string(msg))
	_ = s.Write(msg)
}

// disconnectionHandler is called when a websocket connection is closed.
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
