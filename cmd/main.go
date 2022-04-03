package main

import (
	"log"

	"grpc_go_blog/services/server"
)

func main() {
	if err := server.RunServer(); err != nil {
		log.Println("failed to run server")
		return
	}
}