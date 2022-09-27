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

func HallChatHandler(s *melody.Session, message model.Message) {
	fmt.Println("HallChat")
	msg, err := model.NewMessage(common.HallChat, message.Data)
	if err != nil {
		log.Error(err)
		return
	}

	SendToClient(s, msg)
	BroadcastOthers(s, msg)
}

func NewRoomHandler(s *melody.Session, message model.Message) {
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

	roomId := uuid.NewString()
	room := model.Room{
		ID:             roomId,
		DialogMessages: []model.DialogMessage{},
		Players:        []model.Player{player.(model.Player)},
		Host:           player.(model.Player),
	}
	roomMap.Store(roomId, room)
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
}
