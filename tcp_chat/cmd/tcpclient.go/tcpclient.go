package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"tcp_chat/config"
)

func main() {
	config, err := config.LoadServerConfig(config.ServerConfigPath)
	if err != nil {
		log.Fatalf("Failed load server config: %v", err)
	}

	// Connect to the TCP chat server
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", config.Port))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	fmt.Printf("Connected to chat server on localhost:%d\n", config.Port)

	pid := os.Getpid()
	text := fmt.Sprintf("Hi everyone, I am %d, waiting for the broadcast\n", pid)

	// Send the message to the server
	_, err = conn.Write([]byte(text))
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println("Message sent. Waiting for broadcast...")

	// Read one response from the server (broadcasted message)
	serverReader := bufio.NewReader(conn)
	resp, err := serverReader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	fmt.Printf("Received from server: %s", resp)
}
