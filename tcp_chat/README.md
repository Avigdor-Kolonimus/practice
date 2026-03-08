# TCP Chat

A simple TCP chat server and client in Go. Clients connect to the server, send messages, and receive broadcasts of all messages from other connected clients.

## Features

- TCP server that accepts multiple concurrent connections
- Broadcasts every message to all connected clients
- Configurable port via JSON config
- Graceful shutdown on SIGINT/SIGTERM

## Prerequisites

- Go 1.26+

## Configuration

Server and client read the port from `config/practice/conn.json`:

```json
{
    "port": 3000
}
```

## Running

**Start the server:**

```bash
go run ./cmd/tcpserver
```

**Run the client:**

```bash
go run ./cmd/tcpclient
```

The client connects, sends a greeting with its process ID, waits for the broadcast, then exits.

**Connect with telnet** (for interactive use):

```bash
telnet localhost 3000
```

## Docker

**Build and run with Docker Compose:**

```bash
docker-compose up --build
```

This starts the server (port 3000) and runs the client once. The server keeps running. To run the client again: `docker-compose run tcpclient`.

**Run server only:**

```bash
docker build -f Dockerfile.server -t tcp-chat-server .
docker run -p 3000:3000 tcp-chat-server
```

**Run client** (connects to server at `localhost:3000` when run locally):

```bash
docker build -f Dockerfile.client -t tcp-chat-client .
docker run --network host tcp-chat-client
```

**Run client against Docker server** (set host for container networking):

```bash
docker run -e TCP_CHAT_SERVER_HOST=host.docker.internal tcp-chat-client
```

On Linux, use `--add-host=host.docker.internal:host-gateway` to reach the host.

## Project Structure

```
tcp_chat/
├── cmd/
│   ├── tcpserver/     # Server entry point
│   └── tcpclient/     # Client entry point
├── config/
├── Dockerfile.server
├── Dockerfile.client
├── docker-compose.yml
│   ├── configurator.go
│   └── practice/
│       └── conn.json  # Port configuration
├── message/           # Message type
├── peer/              # Peer/connection metadata
├── server/            # Server logic
├── go.mod
└── README.md
```
