package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello World")
	fmt.Fprintln(os.Stdout, "Hello World")

	fw := NewFileWriter("data.txt")

	fmt.Fprintln(fw, "Hello Golang Folks, Happy Learning!")

}

type FileWriter struct {
	FileName string
}

func NewFileWriter(fn string) *FileWriter {
	return &FileWriter{fn}
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	if fw == nil {
		return 0, fmt.Errorf("nil FileWriter")
	}

	if fw.FileName == "" {
		return 0, errors.New("invalid file name")
	}

	if f, err := os.OpenFile(fw.FileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644); err != nil {
		return 0, err
	} else {
		return f.Write(p)
	}

}

// type Writer interface {
//     Write(p []byte) (n int, err error)
// }

// 0 4+2+1 4+2+1 4+2+1 -> 0777

// 0 4+2 4 4 -> 0644

// Interfaces are contracts
// Interface does not contain any kind of implmentation
// Interfaces are just definitions

// in go no need explicitely say that the type is implementating the interface
// if a type implements all the definitions of an interface, the type auto implement the inteface
