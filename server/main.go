package main

import (
	"encoding/json"
	"fmt"
	"server/common"
	"server/model"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	log "github.com/sirupsen/logrus"
)

var m *melody.Melody

// sessionMap is a map of session ids to melody sessions.
var sessionMap sync.Map
var roomMap sync.Map
var playerMap sync.Map
var playerCounter int

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

// messageHandler is called when a message is received from a client.
func messageHandler(s *melody.Session, msg []byte) {
	var message model.Message
	if err := json.Unmarshal(msg, &message); err != nil {
		errMsg, _ := model.NewErrorMessage(err)
		s.Write(errMsg)
	}

	switch message.Code {
	case common.HallChatRequest:
		HallChatRequestHandler(s, message)
	case common.NewRoomRequest:
		NewRoomRequestHandler(s, message)
	case common.AllRoomsRequest:
		AllRoomsRequestHandler(s)
	case common.HallPlayersRequest:
		HallPlayersRequestHandler(s)
	case common.JoinRoomRequest:
		JoinRoomRequestHandler(s, message)
	default:
		errMsg, _ := model.NewErrorMessage(fmt.Errorf("unknown message code: %d", message.Code))
		_ = s.Write(errMsg)
	}
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
