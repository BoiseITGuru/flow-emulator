package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/yaacov/observer/observer"
)

type WsServer struct {
	wsServer   *http.Server
	eventsChan *observer.Observer
	logger     *logrus.Logger
}

var o *observer.Observer

func NewWsServer(eventsChan *observer.Observer, logger *logrus.Logger) *WsServer {
	wsServer := &http.Server{
		Addr:    ":5050",
		Handler: http.HandlerFunc(wsEndpoint),
	}

	o = eventsChan

	return &WsServer{wsServer: wsServer, eventsChan: eventsChan, logger: logger}
}

func (ws *WsServer) Start() error {
	ws.logger.Info("🌱  WebSocket server connected, listening for connections from Emerald City Playground")
	err := ws.wsServer.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (ws *WsServer) Stop() {
	ws.wsServer.Shutdown(context.Background())
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	o.Emit("ide-connected")

	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				o.Emit("ide-disconnected")
				return
			}

			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				o.Emit("ide-disconnected")
				return
			}
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}
