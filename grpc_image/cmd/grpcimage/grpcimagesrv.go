package main

import (
	"log/slog"

	grpcing "grpc_image/handling"
)

const (
	port = ":8087"
)

func main() {
	imageServer := grpcing.NewImageServer()
	grpcServer := grpcing.NewGrpcServer()

	err := grpcServer.GrpcServeServer(imageServer, port)
	if err != nil {
		slog.Warn("Server shutdown with error", "error", err.Error())
	}
}
