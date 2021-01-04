package main

import (
	"fmt"
)

// Writer is an interface
type Writer interface {
	Write([]byte) (int, error)
}

// Embedded is an interface with embedded interface Writer
type Embedded interface {
	Writer
	Close()
}

// ConsoleWriter implements Writer interface
type ConsoleWriter struct {
	data []byte
}

// NetworkWriter implements Writer interface
type NetworkWriter struct {
	data []byte
}

func (w *ConsoleWriter) Write(data []byte) (int, error) {
	fmt.Println("This is consoleWrite")
	return 0, nil
}

func (w *NetworkWriter) Write(data []byte) (int, error) {
	fmt.Println("This is networkWrite")
	return 0, nil
}

func (w *NetworkWriter) Close() {
	fmt.Println("This is networkWrite Close")
}

func main() {
	var w1 Writer = &ConsoleWriter{} // We dont have to define the type to be *Writer
	var w2 Embedded = &NetworkWriter{}

	w1.Write([]byte{1, 2, 3})
	w2.Write([]byte{1, 2, 3})
	w2.Close()

	var t interface{} = &NetworkWriter{} // We can assign empty interface for any data type
	nw := t.(*NetworkWriter)             // Type conversion via an empty interface assignment

	nw.Write([]byte{1, 2, 3})

}
