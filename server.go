package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

type ChatServer struct {
	messages []string
	mu       sync.Mutex
}

type Message struct {
	Name    string
	Content string
}

type Response struct {
	History []string
}

// Remote procedure: Add message and return history
func (c *ChatServer) SendMessage(msg Message, res *Response) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if msg.Content == "" {
		return errors.New("empty message")
	}

	// append new message to history
	c.messages = append(c.messages, fmt.Sprintf("%s: %s", msg.Name, msg.Content))

	// return history
	res.History = c.messages
	return nil
}

func main() {
	chat := new(ChatServer)
	rpc.Register(chat)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	fmt.Println("Chat server running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
