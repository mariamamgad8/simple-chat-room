# simple-chat-room
📖 Overview

This project implements a simple chatroom using Go RPC (Remote Procedure Call).
Multiple clients can connect to a central server, send messages, and retrieve full chat history.
Every message sent by a client is broadcasted to all clients through shared history.

🛠 Components

The system consists of two main Go files:

File	Role
server.go	Stores chat messages & responds to RPC calls
client.go	Connects to server, sends messages, retrieves chat history
📌 How It Works
✅ Server (server.go)

Starts an RPC server listening on port 1234

Holds a slice of all chat messages: []string

Exposes a remote procedure named SendMessage

When a client sends a message:

The server stores it in the history list

Returns the full chat history to the client

Key structs used:

type Message struct {
	Name    string
	Content string
}

type Response struct {
	History []string
}


The SendMessage RPC method:

func (c *ChatServer) SendMessage(msg Message, res *Response) error {
    c.messages = append(c.messages, fmt.Sprintf("%s: %s", msg.Name, msg.Content))
    res.History = c.messages
    return nil
}

✅ Client (client.go)

Connects to the RPC server using rpc.Dial("tcp", "localhost:1234")

Reads the user's name once

Continuously reads messages using bufio.NewReader (so it accepts spaces)

Calls the server’s SendMessage procedure

After each message, prints full chat history returned from the server

Exits when the user types exit

Example RPC call:

msgObj := Message{Name: name, Content: msg}
client.Call("ChatServer.SendMessage", msgObj, &reply)

🧠 Reason for Using bufio.NewReader Instead of fmt.Scan

fmt.Scan reads input token by token, meaning it stops at spaces.
To allow users to send full sentences, we use:

reader.ReadString('\n')


This reads the entire line including spaces.

▶️ How to Run

Open three terminals in the project directory.

Start the server:

go run server.go


Start first client:

go run client.go


Start second client:

go run client.go


Both clients will now share the same chat history.

✅ Features Implemented

Multiple clients can join the chat

Full chat history is always shown

Messages are stored on server (centralized storage)

Safe concurrent access using mutex in server

📌 Example Output

Client 1

Enter your name: Mariam
Enter message: Hi
--- Chat History ---
Mariam: Hi


Client 2

Enter your name: Ahmed
Enter message: Hello
--- Chat History ---
Mariam: Hi
Ahmed: Hello

🏁 Conclusion

This assignment demonstrates how Go’s RPC mechanism can be used to implement a simple real-time chat system.
The server stores messages, and clients retrieve history through remote procedure calls.
