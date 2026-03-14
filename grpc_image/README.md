## gRPC Image Processing Service

This project is a Go gRPC service that accepts images, applies a watermark, resizes them to requested dimensions, and returns the storage paths of the processed images.

### Requirements

- Go 1.26+
- `protoc` with:
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`
- Docker / Docker Compose (optional, for containerized runs)

### Proto generation

From the project root:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  -I . imageproto/image.proto
```

### Run server locally

```bash
go run ./cmd/grpcimage
```

The server listens on `:8087`.

### Test client (Go)

`cmd/grpcclient/testClient.go` sends a PNG image (`test.png` in the project root) to the server:

```bash
go run ./cmd/grpcclient
```

Processed images are written under `./download/<year>/<month>/<day>/<uuid>/img/`.

### Docker

Build and run via Docker Compose:

```bash
docker compose up --build
```

This:

- Builds the gRPC server image using Go 1.26.
- Exposes the server on `localhost:8087`.
- Mounts `./download` and `./source` into the container.

To stop:

```bash
docker compose down
```

