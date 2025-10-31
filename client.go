package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	serverAddress := "localhost:12345"
	client, err := rpc.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer client.Close()

	// We need the reader to get the name and messages
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read name: %v", err)
	}
	name = strings.TrimSpace(name)

	fmt.Printf("Welcome %s! You've joined the chat. Type a message to see the chat history.\n", name)

	for {
		var history []string

		fmt.Print("Enter message (or 'exit' to quit): ")
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Read error: %v", err)
			continue
		}

		message = strings.TrimSpace(message)

		if message == "exit" {
			break
		}
		if message == "" {
			continue
		}

		formattedMessage := fmt.Sprintf("%s: %s", name, message)

		err = client.Call("ChatRoom.SendMessage", formattedMessage, &history)
		if err != nil {
			log.Printf("RPC call failed: %v", err)
			log.Println("Connection lost. Exiting.")
			break
		}

		fmt.Println("\n--- Chat History ---")
		for _, msg := range history {
			fmt.Println(msg)
		}
		fmt.Println("--------------------")
	}

	fmt.Println("Client shut down.")
}
