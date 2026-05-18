package main

import (
	"context"
	"fmt"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	ctx := context.Background()
	transport := &mcp.SSEClientTransport{Endpoint: "http://localhost:8080"}
	client := mcp.NewClient(&mcp.Implementation{Name: "test-client", Version: "1.0.0"}, nil)
	session, err := client.Connect(ctx, transport, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	result, err := session.CallTool(ctx, &mcp.CallToolParams{
		Name: "ask_weather",
		Arguments: map[string]any{
			"city": "Sofia",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	text := result.Content[0].(*mcp.TextContent).Text
	fmt.Println(text)
}
