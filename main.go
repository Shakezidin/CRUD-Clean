package main

import (
	"github.com/shakezidin/di"
)

func main() {
	server := di.Init()
	server.StartServer()
}
