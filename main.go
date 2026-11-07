// Sample app for the ground-truth tests. It pins a vulnerable dependency
// version ON PURPOSE. Never deploy it.
//
// GHSA-3xh2-74w9-5vxm says "integer overflow in gorilla/websocket" and names
// no function, so the advisory gives nothing to match against.
package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{ReadBufferSize: 1024}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	_, msg, _ := conn.ReadMessage()
	_ = conn.WriteMessage(websocket.TextMessage, msg)
}

func main() {
	http.HandleFunc("/ws", handler)
	_ = http.ListenAndServe("127.0.0.1:8080", nil)
}
