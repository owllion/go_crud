package websocket

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnection(c *gin.Context) {
	// 升級 initial GET request 到 websocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Upgrade Error:", err)
		return
	}
	defer ws.Close()

	fmt.Println("Client Connected")

	// 無限循環，持續監聽來自客戶端的消息
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Read Error:", err)
			return
		}
		// 打印接收到的消息
		fmt.Println("Received:", string(p))

		// 將消息發送回客戶端
		if err := ws.WriteMessage(messageType, p); err != nil {
			fmt.Println("Write Error:", err)
			return
		}
	}
}
