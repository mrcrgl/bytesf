package main

import (
	"fmt"

	"github.com/mrcrgl/bytesf"
)

func main() {
	rb := bytesf.NewBufferPool(64, 256)

	// Allocate buffer
	b := rb.Allocate()

	b.WriteString("Hello World!")
	fmt.Printf("% x", b.Bytes())

	// Release ownership of buffer
	rb.Release(b)
}
