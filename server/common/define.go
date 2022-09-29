package common

type MessageCode int

const (
	Error               = iota
	Success             = 1
	HallChatRequest     = 2
	NewRoomRequest      = 3
	DeleteRoomRequest   = 4
	JoinRoomRequest     = 5
	LeaveRoomRequest    = 6
	RoomChatRequest     = 7
	AllRoomsRequest     = 8
	HallPlayersRequest  = 9
	HallChatResponse    = 202
	NewRoomResponse     = 203
	DeleteRoomResponse  = 204
	JoinRoomResponse    = 205
	LeaveRoomResponse   = 206
	RoomChatResponse    = 207
	AllRoomsResponse    = 208
	HallPlayersResponse = 209
)
