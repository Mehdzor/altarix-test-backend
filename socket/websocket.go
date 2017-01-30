package socket

import (
	"altarix_test/model"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

var ActiveSocket *WebSocket

type WebSocket struct {
	hub *hub
}

func Run() {
	h := newHub()
	s := WebSocket{h}
	ActiveSocket = &s
	h.Run()
}

func (ws *WebSocket)Serve(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: ws.hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client
	client.writePump()
}


func (ws *WebSocket)Send(e model.Event) {
	b, err := json.Marshal(e)
	if err != nil {
		log.Error(err)
		return
	}
	ws.hub.broadcast <- b
}