package main

import (
    "fmt"
    "github.com/gorilla/websocket"
)

func main() {
    fmt.Println("Gorilla WebSocket Version Demo")
    // Demonstrate usage of the vulnerable dependency
    _ = websocket.IsCloseError
}
