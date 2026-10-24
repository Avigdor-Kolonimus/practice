package mobiledev

import (
	"bufio"
	"encoding/binary"
	"errors"
	"os"
	"strings"
)

var (
	ErrNotRingBuffer = errors.New("not a ring buffer")
	ErrDataLoss      = errors.New("data loss")
)

func solution(f *os.File) error {
	st, err := f.Stat()
	if err != nil {
		return ErrNotRingBuffer
	}

	fileSize := st.Size()

	var buf [4]byte
	blocksInRingTotalSize := int64(0)
	blockIndex := int64(0)

	for blocksInRingTotalSize < fileSize {
		dataSize, err := ReadInt64At(f, buf, blockIndex)
		if err != nil {
			return ErrNotRingBuffer
		}

		nextBlockIndex, err := ReadInt64At(f, buf, blockIndex+4+dataSize)
		if err != nil {
			return ErrNotRingBuffer
		}

		blocksInRingTotalSize += dataSize + 8
		blockIndex = nextBlockIndex

		if blockIndex == 0 {
			break
		}
	}

	if blockIndex != 0 {
		return ErrNotRingBuffer
	} else if fileSize != blocksInRingTotalSize {
		return ErrDataLoss
	}

	return nil
}

func ReadInt64At(f *os.File, buf [4]byte, at int64) (int64, error) {
	if _, err := f.ReadAt(buf[:], at); err != nil {
		return 0, err
	}

	dataSize := binary.LittleEndian.Uint32(buf[:])
	return int64(dataSize), nil
}

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/tos-23
// Tos23 - problem 16
func Tos23() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	filename := strings.TrimSpace(line)

	f, err := os.Open(filename)
	if err != nil {
		writer.WriteString(ErrNotRingBuffer.Error())
		writer.WriteByte('\n')
		return
	}
	defer f.Close()

	err = solution(f)
	if err != nil {
		writer.WriteString(err.Error())
		writer.WriteByte('\n')
	} else {
		writer.WriteString("OK")
		writer.WriteByte('\n')
	}
}
