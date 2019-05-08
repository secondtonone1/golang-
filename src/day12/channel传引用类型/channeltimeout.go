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
	//此时data 变为aaa，但是地址没有变化。
	//而之前写入chan中写入的大小为10的字节切片bbbbbbbbbb，由于切片每隔元素是地址引用，
	//所以此时chan中此时数据变为aaabbbbbbb
	//再次写入chan中，此时chan中存在两个元素一个aaabbbbbbb,一个aaa
	ch <- data

	time.Sleep(time.Second * 5)
}
