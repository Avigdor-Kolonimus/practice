package examples

import (
	"bytes"
	"fmt"
)

func PeekBuffer() {
	buf := bytes.NewBufferString("Hello, Go 1.26!")

	peek, _ := buf.Peek(5)
	fmt.Printf("Peek: %s\n", peek)
	fmt.Printf("Len: %d\n", buf.Len())
	if len(peek) < 100 {
		fmt.Printf("Peek < 100\n")
	}

	data := make([]byte, 5)
	buf.Read(data)
	fmt.Printf("Read: %s\n", data)
	fmt.Printf("Len: %d\n", buf.Len())

	buf.Write([]byte("more data"))
	fmt.Printf("Peek: %s\n", peek)
}
