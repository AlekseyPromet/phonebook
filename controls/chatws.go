package controls

import (
	ws "github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

//Message for chat
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

var (
	clients   = make(map[*ws.Conn]bool)
	broadcast = make(chan Message)
	upgrader  = ws.Upgrader{}
)

//ChatHandler func
func ChatHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}
