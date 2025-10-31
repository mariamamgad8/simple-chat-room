package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

type ChatRoom struct {
	mutex    sync.Mutex
	messages []string
}

// SendMessage is the exported RPC method.
func (cr *ChatRoom) SendMessage(message string, history *[]string) error {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()

	cr.messages = append(cr.messages, message)

	fmt.Println(message)

	// Copy the full history to the reply
	*history = make([]string, len(cr.messages))
	copy(*history, cr.messages)

	return nil
}

func main() {
	chatRoom := new(ChatRoom)
	err := rpc.Register(chatRoom)
	if err != nil {
		log.Fatalf("Failed to register RPC: %v", err)
	}

	port := ":12345"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}
	defer listener.Close()

	// This is the only thing the server logs on startup
	fmt.Printf("Chat server running on port %s\n", port)

	// Main accept loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}

		go rpc.ServeConn(conn)
	}
}
