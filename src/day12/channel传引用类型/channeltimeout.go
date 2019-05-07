package main

import "fmt"
import "time"

func main() {
	ch := make(chan []byte, 10)
	go func() {
		for {
			select {
			case data := <-ch:
				fmt.Println(string(data))
			}
		}
	}()
	data := make([]byte, 0, 32)
	data = append(data, []byte("bbbbbbbbbb")...)
	ch <- data

	//fmt.Printf("%p\n", data)
	data = data[:0]
	//fmt.Printf("%p\n", data)

	data = append(data, []byte("aaa")...)
	ch <- data

	time.Sleep(time.Second * 5)
}
