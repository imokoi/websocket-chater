package model

import (
	"encoding/json"
	"server/common"
)

type Message struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewMessage(code int, data interface{}) ([]byte, error) {
	return json.Marshal(Message{
		Code: code,
		Data: data,
	})
}

func NewErrorMessage(err error) ([]byte, error) {
	return json.Marshal(Message{
		Code: common.Error,
		Data: err,
	})
}
