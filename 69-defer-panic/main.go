package main

import "os"

func main() {

	defer println("Hello World, end of Main")

	_, err := Write("data.txt", "Hello World, How are you doing")
	if err != nil {
		println(err.Error())
	} else {
		println("data successfully stored in the file")
	}

	println("Hello World")

}

func Write(filename string, data string) (int, error) {
	defer func() {
		println("Write function is called , whether panic or not I am getting executed")
	}()
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err.Error()) // user defined panic, custom panic , when ever there is a panic before panic,
		//return 0, err
	}

	defer func() {
		println("There is a panic above, would I be executed?") // No
	}()
	defer f.Close()
	//f.Close()
	return f.Write([]byte(data))
	//return n, err
}
