package main

import (
	"log"
	"fmt"
	
	"tcp_chat/server"
	"tcp_chat/config"
)

func main() {
	config, err := config.LoadServerConfig(config.ServerConfigPath)
	if err != nil {
		log.Fatalf("Failed load server config: %v", err)
	}

	server := server.NewServer(fmt.Sprintf(":%d", config.Port))
	server.Start()
}