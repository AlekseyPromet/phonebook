package controls

import (
	"fmt"

	ws "github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader = ws.Upgrader{}
)

//SearchHandl for search number
func SearchHandl(ctx echo.Context) error {
	wsc, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	defer wsc.Close()

	for {
		// Write
		err := wsc.WriteMessage(ws.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			ctx.Logger().Error(err)
		}

		// Read
		_, msg, err := wsc.ReadMessage()
		if err != nil {
			ctx.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}
