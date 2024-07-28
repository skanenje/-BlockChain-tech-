package main

import (
    "fmt"
    "swapBlock/internal/server"
)

func main() {
    fmt.Println("Starting blockchain server...")
    s := server.NewServer("node1")
    s.Start()
}