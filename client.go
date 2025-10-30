package main

import (
    "bufio"
    "fmt"
    "net/rpc"
    "os"
)

type Message struct {
    Name    string
    Content string
}

type Response struct {
    History []string
}

func main() {
    client, err := rpc.Dial("tcp", "localhost:1234")
    if err != nil {
        fmt.Println("❌ Error connecting to server:", err)
        return
    }

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your name: ")
    name, _ := reader.ReadString('\n')
    name = name[:len(name)-1]

    fmt.Printf("Welcome %s! You've joined the chat. Type message (or 'exit' to quit):\n", name)

    for {
        fmt.Print("Enter message: ")
        msg, _ := reader.ReadString('\n')
        msg = msg[:len(msg)-1]

        if msg == "exit" {
            fmt.Println("Leaving chat...")
            return
        }

        msgObj := Message{Name: name, Content: msg}
        var reply Response

        err = client.Call("ChatServer.SendMessage", msgObj, &reply)
        if err != nil {
            fmt.Println("❌ Server Error:", err)
            return
        }

        fmt.Println("\n--- Chat History ---")
        for _, m := range reply.History {
            fmt.Println(m)
        }
        fmt.Println("---------------------\n")
    }
}
