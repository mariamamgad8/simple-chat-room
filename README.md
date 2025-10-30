# 💬 Simple Chatroom Using Go RPC

A simple chatroom application built using **Go RPC (Remote Procedure Call)**.  
Multiple clients can connect to the server, send messages, and retrieve the full chat history.

---

## 🚀 Features

✅ Multiple clients can join the chat  
✅ Server stores messages in memory  
✅ Client retrieves full chat history after every message  
✅ Uses `bufio.Reader` so messages can contain spaces  
✅ Exits gracefully using `exit` or `Ctrl + C`  

---

## 🛠 Technologies Used

- Go (Golang)
- net/rpc package
- TCP networking

---

## ⚙️ How It Works

### ✅ Server (`server.go`)

The server:
- Runs an RPC server on port **1234**
- Stores all messages in a `[]string`
- Exposes a remote function `SendMessage`
- Returns the full chat history after each message

Message and response models:

```go
type Message struct {
    Name    string
    Content string
}

type Response struct {
    History []string
}
```

### ✅ Client (`client.go`)

The client:

- Connects to server via `rpc.Dial("tcp", "localhost:1234")`
- Uses `bufio.Reader` to allow messages with spaces
- Sends user input as Message struct
- Displays full conversation history
- Stops when user types exit

---


