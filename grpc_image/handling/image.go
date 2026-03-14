package handling

import (
	"errors"
	"io"
	"strconv"

	pb "grpc_image/imageproto"
)

const (
	defaultWatermark = "source/watermark.png"
)

type ImageServer struct {
	pb.ImageServiceServer
}

func NewImageServer() ImageServer {
	imgs := ImageServer{}

	return imgs
}

type OriginalImage struct {
	Path      string
	Lenght    []int32
	Width     []int32
	Format    string
	Folder    string
	Watermark string
	UUID      string
}

func (imgs ImageServer) DownloadImages(stream pb.ImageService_DownloadImagesServer) error {
	var images []*pb.DownloadImagesRequest
	for {
		image, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return stream.SendAndClose(&pb.DownloadImagesResponse{
				Error: err.Error(),
			})
		}

		images = append(images, image)
	}

	var paths []OriginalImage
	for i := range images {
		if len(images[i].Info.Height) != len(images[i].Info.Width) {
			return stream.SendAndClose(&pb.DownloadImagesResponse{
				Error: "different len of lenght and width for picture " + strconv.Itoa(i),
			})
		}
		if images[i].Info.Watermark == "" {
			images[i].Info.Watermark = defaultWatermark
		}
	}

	paths = saveSourceFiles(images)
	if len(images) == 0 {
		return stream.SendAndClose(&pb.DownloadImagesResponse{
			Error: "no images in request",
		})
	}

	err := watermark(paths)
	if err != nil {
		return stream.SendAndClose(&pb.DownloadImagesResponse{
			Error: errors.New("path for watermark is invalid").Error(),
		})
	}

	uploadPath := resizeAndSave(paths)

	res := &pb.DownloadImagesResponse{
		StoragePath: uploadPath,
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return stream.SendAndClose(&pb.DownloadImagesResponse{
			Error: "error when receive an responce",
		})
	}

	return nil
}
