package main

import (
	"MESSAGE-SERVICE/server"
	"fmt"
)

func main() {
	// server.ServerRun()
	// server.HttpServerRun()
	server.HttpAndGrpc()
	fmt.Printf("***********")
}
