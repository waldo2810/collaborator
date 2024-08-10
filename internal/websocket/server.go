package websocket

import (
	"fmt"
	"io"
	"sync"

	"golang.org/x/net/websocket"
)

type Server struct {
	clients map[*websocket.Conn]bool
	mutex   sync.Mutex
}

func NewServer() *Server {
	return &Server{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Error. Client probably disconnected.")
				break
			}
			fmt.Println("Error reading data from ws frame.")
			continue
		}
		msg := buf[:n]
		fmt.Printf("Received message\nClient: %q\nMessage: %q", ws.RemoteAddr(), msg)
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	fmt.Printf("New connection from address %q.", ws.RemoteAddr())

	s.mutex.Lock()
	s.clients[ws] = true
	s.mutex.Unlock()

	s.readLoop(ws)
}
