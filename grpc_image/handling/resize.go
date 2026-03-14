package handling

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"log/slog"
	"os"
	"strconv"
	"sync"

	"github.com/nfnt/resize"
)

func resizeAndSave(paths []OriginalImage) []string {
	var wg sync.WaitGroup
	var mu sync.Mutex

	uploadPaths := make([]string, 0)
	for i := range paths {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			uploadPath := resizeImage(paths[i])
			if len(uploadPath) != 0 {
				mu.Lock()
				uploadPaths = append(uploadPaths, uploadPath...)
				mu.Unlock()
			} else {
				mu.Lock()
				uploadPaths = append(uploadPaths, paths[i].Path)
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()
	return uploadPaths
}

func resizeImage(path OriginalImage) []string {
	var mu sync.Mutex

	for i := range path.Lenght {
		if path.Lenght[i] == 0 || path.Width[i] == 0 {
			return []string{}
		}
	}

	if len(path.Lenght) != len(path.Width) {
		return []string{}
	}

	var uploadPaths []string
	switch path.Format {
	case "png":
		for i := range path.Lenght {
			imgIn, err := os.Open(path.Path)
			if err != nil {
				slog.Error("failed to open PNG file", "error", err.Error())

				return []string{}
			}

			imgPng, err := png.Decode(imgIn)
			if err != nil {
				slog.Error("failed to decode PNG", "error", err.Error())

				return []string{}
			}

			err = imgIn.Close()
			if err != nil {
				slog.Error("failed to close PNG file", "error", err.Error())

				return []string{}
			}

			imgPng = resize.Resize(uint(path.Lenght[i]), uint(path.Width[i]), imgPng, resize.Bilinear)
			upPath := path.Folder + path.UUID + "_" + strconv.FormatUint(uint64(path.Lenght[i]), 10) + "x" + strconv.FormatUint(uint64(path.Width[i]), 10) + "." + path.Format
			buf := new(bytes.Buffer)

			err = png.Encode(buf, imgPng)
			if err != nil {
				slog.Error("failed to encode PNG", "error", err.Error())

				return []string{}
			}

			imgSave := buf.Bytes()

			err = os.WriteFile(upPath, imgSave, 0666)
			if err != nil {
				slog.Error("failed to save PNG file", "error", err.Error())

				return []string{}
			}

			mu.Lock()
			uploadPaths = append(uploadPaths, upPath)
			mu.Unlock()
		}
	case "jpg", "jpeg":
		for i := range path.Lenght {
			imgIn, err := os.Open(path.Path)
			if err != nil {
				slog.Error("failed to open JPEG file", "error", err.Error())

				return []string{}
			}

			imgJpeg, err := jpeg.Decode(imgIn)
			if err != nil {
				slog.Error("failed to decode JPEG", "error", err.Error())

				return []string{}
			}

			err = imgIn.Close()
			if err != nil {
				slog.Error("failed to close JPEG file", "error", err.Error())

				return []string{}
			}

			imgJpeg = resize.Resize(uint(path.Lenght[i]), uint(path.Width[i]), imgJpeg, resize.Bilinear)
			upPath := path.Folder + path.UUID + "_" + strconv.FormatUint(uint64(path.Lenght[i]), 10) + "x" + strconv.FormatUint(uint64(path.Width[i]), 10) + "." + path.Format
			buf := new(bytes.Buffer)

			err = jpeg.Encode(buf, imgJpeg, &jpeg.Options{Quality: 100})
			if err != nil {
				slog.Error("failed to encode JPEG", "error", err.Error())
				return []string{}
			}

			imgSave := buf.Bytes()

			err = os.WriteFile(upPath, imgSave, 0666)
			if err != nil {
				slog.Error("failed to save JPEG file", "error", err.Error())
				return []string{}
			}

			mu.Lock()
			uploadPaths = append(uploadPaths, upPath)
			mu.Unlock()
		}
	default:
		slog.Error("unknown file format, please provide a file with extension png or jpg")
	}

	return uploadPaths
}
