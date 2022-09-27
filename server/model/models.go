package model

type DialogMessage struct {
	Time    string `json:"time"`
	From    string `json:"from"`
	Content string `json:"content"`
}

type Player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Room struct {
	ID             string          `json:"id"`
	DialogMessages []DialogMessage `json:"dialogMessages"`
	Players        []Player        `json:"players"`
	Host           Player          `json:"host"`
}
