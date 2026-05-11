package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"sync"
)

func main() {

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() { // func1 -> P1
		println("Running a G1")
		wg.Add(1)
		go func() { //func1.1 P1
			wg.Add(1)
			go func() {
				println("Running G2") // P1
				wg.Done()
			}()
			wg.Add(1)
			go func() { // func1.1.1 P1

				for i := 1; i <= 10; i++ {

					if i%2 == 0 {
						wg.Add(1)
						go func() {
							println("even-->", i) // func1.1.1.i P1
							wg.Done()
						}()
					}
				}
				wg.Done()
			}()
			wg.Done()
		}()
		wg.Done()
	}()

	wg.Add(1)
	go func() { // This is an IO operation
		f, err := os.OpenFile("data.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			println(err.Error())
			wg.Done()
			runtime.Goexit()
		}
		defer f.Close()
		_, err = f.Write([]byte("Hello Wrold"))
		if err != nil {
			println(err.Error())
			wg.Done()
			runtime.Goexit()
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() { // network IO

		url := "https://httpbin.org/get"
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			println(err.Error())
			wg.Done()
			runtime.Goexit()
		}

		req.Header.Set("accept", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			println(err.Error())
			wg.Done()
			runtime.Goexit()
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			println(err.Error())
			wg.Done()
			runtime.Goexit()
		}

		fmt.Print("Status:", resp.Status)
		fmt.Println(string(body))
		wg.Done()

	}()

	wg.Wait()
}
