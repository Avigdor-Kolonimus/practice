package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"

	pb "grpc_image/imageproto"
)

func main() {
	conn, err := grpc.Dial("localhost:8087", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewImageServiceClient(conn)

	stream, err := client.DownloadImages(context.Background())
	if err != nil {
		log.Fatalf("failed to create stream: %v", err)
	}

	imgBytes, err := os.ReadFile("source/test_image.png")
	if err != nil {
		log.Fatalf("failed to read image: %v", err)
	}

	req := &pb.DownloadImagesRequest{
		Info: &pb.ImageInfo{
			Compress:  "medium",
			Watermark: "",
			Format:    "png",
			Width:     []int32{800},
			Height:    []int32{600},
		},
		Image: imgBytes,
	}

	if err := stream.Send(req); err != nil {
		log.Fatalf("failed to send request: %v", err)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive response: %v", err)
	}

	fmt.Printf("storage paths: %+v, error: %q\n", resp.StoragePath, resp.Error)
}
