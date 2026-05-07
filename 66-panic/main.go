package main

import "os"

func main() {

	_, err := Write("data.txt", "Hello World, How are you doing")
	if err != nil {
		println(err.Error())
	} else {
		println("data successfully stored in the file")
	}

	println("Hello World")

}

func Write(filename string, data string) (int, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err.Error()) // user defined panic, custom panic
		//return 0, err
	}
	//f.Close()
	defer f.Close()
	return f.Write([]byte(data))
	//return n, err
}
