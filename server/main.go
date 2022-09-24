package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/olahol/melody"
	log "github.com/sirupsen/logrus"
)

var m *melody.Melody

// sessionMap is a map of session ids to melody sessions.
var sessionMap sync.Map

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	r := gin.Default()
	m = melody.New()
	// add a ping handler with gin
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// add a websocket handler with gin, but we are using melody to handle the requests
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
	_ = s.Write([]byte("welcome to chatting hall"))
	_ = m.BroadcastOthers([]byte("new member joined chatting hall"), s)
}

// messageHandler is called when a message is received from a client.
func messageHandler(s *melody.Session, msg []byte) {
	fmt.Println(string(msg))
	_ = s.Write(msg)
	_ = m.BroadcastOthers([]byte(string(msg)), s)
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
