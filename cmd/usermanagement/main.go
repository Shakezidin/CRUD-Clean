package main

import "github.com/shakezidin/internal/di"

func main() {
	server := di.Init()
	server.StartServer()
}
