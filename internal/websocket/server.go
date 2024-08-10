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

func (s *Server) broadcast(msg []byte) {
	fmt.Println("\nBroadcasting...")
	for ws := range s.clients {
		connected := s.clients[ws]
		if connected {
			_, err := ws.Write(msg)
			if err != nil {
				fmt.Println("Error broadcasting msg", err)
			}
		}
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
		fmt.Printf("\nReceived message from Client: %q. Content: %q.", ws.RemoteAddr(), msg)
		s.broadcast(msg)
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	fmt.Printf("\nNew connection from address %q.", ws.RemoteAddr())

	s.mutex.Lock()
	s.clients[ws] = true
	s.mutex.Unlock()

	s.readLoop(ws)
}
