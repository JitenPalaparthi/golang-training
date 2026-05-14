package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Line struct {
	No   int
	Data string
}

func NewLine(no int, data string) *Line {
	return &Line{no, data}
}

func main() {
	lineCh := make(chan *Line, 10)
	go ReadFileCh("data.txt", lineCh)

	linech1, linech2, linech3 := FanOut(lineCh)
	//_, _, linech3 := FanOut(lineCh)

	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		for line := range linech1 {
			words := strings.Split(line.Data, " ")
			count := 0
			for _, w := range words {
				if len(w) != 0 {
					count++
				}
			}
			fmt.Println("Line No", line.No, "Words Count", count)
		}
		wg.Done()
	}()

	go func() {
		for line := range linech2 {
			// please take care of spaces
			count := 0
			for _, c := range line.Data {
				if c != ' ' {
					count++
				}
			}
			//fmt.Println("Line No", line.No, "Char Count", len(line.Data))
			fmt.Println("Line No", line.No, "Char Count", count)
		}
		wg.Done()
	}()

	go func() {
		for line := range linech3 {
			// please take care of spaces
			count := 0
			for _, c := range line.Data {
				switch c {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					count++
				}
			}
			fmt.Println("Line No", line.No, "Vowel Count:", count)
		}
		wg.Done()
	}()
	wg.Wait()
}

func ReadFileCh(filename string, lineCh chan<- *Line) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		println(err.Error())
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	no := 0

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			close(lineCh)
			break
		}

		if err != nil {
			fmt.Println(err)
			close(lineCh)
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		} else {
			no++
			lineCh <- NewLine(no, line)
		}
	}
}

func FanOut(linech chan *Line) (linech1, linech2, linech3 chan *Line) {
	linech1 = make(chan *Line, 10)
	linech2 = make(chan *Line, 10)
	linech3 = make(chan *Line, 10)

	go func() {
		for line := range linech {
			linech1 <- line
			linech2 <- line
			linech3 <- line
		}
		close(linech1)
		close(linech2)
		close(linech3)

	}()

	return linech1, linech2, linech3
}

func FanOutS(linech chan *Line, chancount uint) []chan *Line {
	chanSlice := make([]chan *Line, chancount)

	go func() {
		for line := range linech {
			for _, ch := range chanSlice {
				ch <- line
			}
		}
		for _, ch := range chanSlice {
			close(ch)
		}
	}()
	return chanSlice
}

// find out number of words
// findout number of chars
