package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func main() {
	addr := ":11050"
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		userId := context.GetHeader("userId")
		context.JSON(http.StatusOK, []string{"hello", userId})
	})
	r.GET("/ws", testWs)
	err := r.Run(addr)
	if err != nil {
		fmt.Println(err)
	}
}

func testWs(c *gin.Context) {
	u := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := u.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(msg) == "test" {
			msg = []byte("websocketTest")
		}
		err = ws.WriteMessage(mt, msg)
		if err != nil {
			break
		}
	}
}
