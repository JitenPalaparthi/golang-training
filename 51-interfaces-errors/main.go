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

	fw1 := NewFileWriter("")

	if _, err := fmt.Fprintln(fw1, "Hello Golang Minds! Learn Well!"); err != nil {
		switch e := err.(type) { // type assertion
		case *FileError:
			fmt.Println("Error Code:", e.Code)
			fmt.Println("Error Message:", e.Message)
		}
	}

	// Same like above but in another wany using As

	if _, err := fmt.Fprintln(fw1, "Hello Golang Minds! Learn Well!"); err != nil {
		fileErr := new(FileError)
		if ok := errors.As(err, &fileErr); ok {
			fmt.Println("Error Code:", fileErr.Code)
			fmt.Println("Error Message:", fileErr.Message)
		}
	}

	// Same like above but in another wany using AsType which is generic implementation

	if _, err := fmt.Fprintln(fw1, "Hello Golang Minds! Learn Well!"); err != nil {
		if fileErr, ok := errors.AsType[*FileError](err); ok {
			fmt.Println("Error Code:", fileErr.Code)
			fmt.Println("Error Message:", fileErr.Message)
		}
	}

}

// type IEmpty interface{} // since no definitions every type implement this interface

// type error interface {
//     Error() string
// }

type FileWriter struct {
	FileName string
}

func NewFileWriter(fn string) *FileWriter {
	return &FileWriter{fn}
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	if fw == nil {
		fe := NewFileError(100, "invalid file writer")
		return 0, fe
	}
	if fw.FileName == "" {
		fe := NewFileError(101, "invalid file name")
		return 0, fe
	}

	if f, err := os.OpenFile(fw.FileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644); err != nil {
		return 0, err
	} else {
		return f.Write(p)
	}
}

type FileError struct {
	Code    int
	Message string
}

func NewFileError(code int, msg string) *FileError {
	return &FileError{code, msg}
}

func (fe *FileError) Error() string {
	return fmt.Sprint("Code:", fe.Code, " ", "Message:", fe.Message)
}
