# mcp_weather

A Go MCP server that returns current weather for a city by name. Transport: **SSE (HTTP)**. Data source: [Open-Meteo](https://open-meteo.com/) (geocoding + forecast).

## Features

- MCP tool `ask_weather` — temperature and wind for a city
- Geocoding cache (repeat requests for the same city are faster)
- Timeouts and cache TTL via JSON config
- Sample Go client and Cursor integration

## Requirements

- Go 1.26+
- Internet access (calls to Open-Meteo)

## Project layout

```text
mcp_weather/
├── cmd/
│   ├── server/          # MCP SSE server
│   └── clienttest/      # sample client
├── config/
│   ├── configuration.go
│   └── weaterserverconfig.json
├── geo/                 # geocoding, weather, MCP handler
├── go.mod
└── README.md
```

## Quick start

From the repository root:

```bash
# 1. Dependencies
go mod download

# 2. Server (port :8080)
go run ./cmd/server

# 3. In another terminal — test client
go run ./cmd/clienttest
```

Expected client output looks like: `Sofia, ... 12.5°C, wind 2.2`.

## Configuration

Default file: `config/weaterserverconfig.json`

```json
{
  "requestTimeout": 8,
  "toolTimeout": 10,
  "cacheTTL": 3600
}
```

| Field | Unit | Description |
|-------|------|-------------|
| `requestTimeout` | seconds | HTTP timeout for Open-Meteo requests |
| `toolTimeout` | seconds | Max duration for `ask_weather` |
| `cacheTTL` | seconds | Geocoding cache TTL |

All fields must be **positive** integers.

Override the config path:

```bash
MCP_WEATHER_CONFIG=/path/to/config.json go run ./cmd/server
```

## MCP tool

| Name | Arguments | Description |
|------|-----------|-------------|
| `ask_weather` | `city` (string) | Current weather for a city |

Example arguments:

```json
{ "city": "Ankara" }
```

## Cursor integration

Add to `.cursor/mcp.json` (or global `~/.cursor/mcp.json`):

```json
{
  "mcpServers": {
    "weather": {
      "url": "http://localhost:8080"
    }
  }
}
```

Start the server, reload MCP in Cursor, then in Agent: *“Use ask_weather for Istanbul”*.

## Protocol (SSE)

The server uses the [MCP SSE transport](https://modelcontextprotocol.io/specification/2024-11-05/basic/transports):

1. **GET** `http://localhost:8080` with `Accept: text/event-stream` — open a session, read `sessionid`
2. **POST** JSON-RPC to `http://localhost:8080?sessionid=...` — `initialize`, `notifications/initialized`, `tools/call`
3. Responses arrive on the SSE stream (`event: message`)

For manual debugging, `go run ./cmd/clienttest` is easier than Postman/Bruno.

## Dependencies

- [github.com/modelcontextprotocol/go-sdk](https://github.com/modelcontextprotocol/go-sdk) — MCP server/client
- Open-Meteo API — no API key required
