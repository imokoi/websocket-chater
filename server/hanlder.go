package main

import (
	"fmt"
	"server/common"
	"server/model"

	"github.com/google/uuid"
	"github.com/olahol/melody"
	log "github.com/sirupsen/logrus"
)

func SendToClient(s *melody.Session, msg []byte) {
	if err := s.Write(msg); err != nil {
		log.Error(err)
	}
}

func BroadcastAll(msg []byte) {
	if err := m.Broadcast(msg); err != nil {
		log.Error(err)
	}
}

func BroadcastOthers(s *melody.Session, msg []byte) {
	if err := m.BroadcastOthers(msg, s); err != nil {
		log.Error(err)
	}
}

func SendSuccess(s *melody.Session) {
	if err := s.Write([]byte("success")); err != nil {
		log.Error(err)
	}
}

func SendFail(s *melody.Session) {
	if err := s.Write([]byte("fail")); err != nil {
		log.Error(err)
	}
}

func HallChatRequestHandler(s *melody.Session, message model.Message) {
	log.Infof("HallChatRequestHandler, message: %v", message)
	msg, err := model.NewMessage(common.HallChatResponse, message.Data)
	if err != nil {
		log.Error(err)
		return
	}

	SendToClient(s, msg)
	BroadcastOthers(s, msg)
}

// NewRoomRequestHandler create a new room and broadcast to all clients
func NewRoomRequestHandler(s *melody.Session, message model.Message) {
	playerId, exist := s.Get("id")
	if !exist {
		log.Error("player id not exist")
		return
	}
	// 1. get current player
	player, ok := playerMap.Load(playerId)
	if !ok {
		log.Error("player not exist")
		return
	}

	// 2. create a new room
	roomId := uuid.NewString()
	room := model.Room{
		ID:             roomId,
		DialogMessages: []model.DialogMessage{},
		Players:        []model.Player{player.(model.Player)},
		Host:           player.(model.Player),
	}
	roomMap.Store(roomId, room)
	var rooms []model.Room
	roomMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		rooms = append(rooms, value.(model.Room))
		return true
	})

	// send a success message to client
	msgToClient, _ := model.NewMessage(common.NewRoomResponse, "create room success")
	SendToClient(s, msgToClient)
	// send all rooms message to others to refresh rooms
	msgToOthers, _ := model.NewMessage(common.AllRoomsResponse, rooms)
	BroadcastAll(msgToOthers)
}

// AllRoomsRequestHandler is called when a client request all rooms
func AllRoomsRequestHandler(s *melody.Session) {
	var rooms []model.Room
	roomMap.Range(func(key, value interface{}) bool {
		rooms = append(rooms, value.(model.Room))
		return true
	})
	msg, _ := model.NewMessage(common.AllRoomsResponse, rooms)
	SendToClient(s, msg)
}

// HallPlayersRequestHandler is called when a client request all players in hall
func HallPlayersRequestHandler(s *melody.Session) {
	var players []model.Player
	playerMap.Range(func(key, value interface{}) bool {
		players = append(players, value.(model.Player))
		return true
	})
	msg, _ := model.NewMessage(common.HallPlayersResponse, players)
	SendToClient(s, msg)
}

// joinRoomHandler is called when a client join a room
func JoinRoomRequestHandler(s *melody.Session, message model.Message) {
	playerId, exist := s.Get("id")
	if !exist {
		log.Error("player id not exist")
		return
	}
	player, ok := playerMap.Load(playerId)
	if !ok {
		log.Error("player not exist")
		return
	}

	roomId := message.Data.(string)
	roomValue, ok := roomMap.Load(roomId)
	if !ok {
		log.Error("room not exist")
		return
	}
	room := roomValue.(model.Room)

	// add player to room
	room.Players = append(room.Players, player.(model.Player))
	roomMap.Store(roomId, room)

	// send this message to other clients
	toClientMsg, _ := model.NewMessage(common.JoinRoomResponse, "join room success")
	SendToClient(s, toClientMsg)

	// send this message to other clients
	for _, p := range room.Players {
		toOthersMsg, _ := model.NewMessage(common.JoinRoomResponse, fmt.Sprintf("%s has joined the room", player.(model.Player).Name))
		sessionValue, ok := sessionMap.Load(p.ID)
		if !ok {
			log.Error("session not exist")
			continue
		}
		session := sessionValue.(*melody.Session)
		SendToClient(session, toOthersMsg)
	}
}

// connectionHandler is called when a new websocket connection is established.
func connectionHandler(s *melody.Session) {
	id := uuid.NewString()
	s.Set("id", id)

	player := model.Player{
		ID:   id,
		Name: fmt.Sprintf("player-%d", playerCounter),
	}
	playerCounter += 1
	playerMap.Store(id, player)
	sessionMap.Store(id, s)
	// send this message to other clients
	toClientMsg, _ := model.NewMessage(common.Success, fmt.Sprintf("welcome %s to server", player.Name))
	SendToClient(s, toClientMsg)

	// send this message to other clients
	toOthersMsg, _ := model.NewMessage(common.Success, fmt.Sprintf("%s has joined the server", player.Name))
	BroadcastOthers(s, toOthersMsg)

	AllRoomsRequestHandler(s)
	HallPlayersRequestHandler(s)
}
