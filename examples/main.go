package main

import (
	"fmt"

	"github.com/mrcrgl/bytesf"
)

func main() {
	rb := bytesf.NewBufferPool(64, 256)

	// Allocate buffer
	b := rb.GetBuffer()

	b.WriteString("Hello World!")
	fmt.Printf("% x", b.Bytes())

	// Release ownership of buffer
	rb.PutBuffer(b)
}
