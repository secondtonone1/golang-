package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	srcfile := "helloworld.txt"
	dstfile := "dstfile.txt"
	data, err := ioutil.ReadFile(srcfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Read error %s\n", err)
		return
	}
	fmt.Println("Has read such data:")
	fmt.Println(string(data))
	wterr := ioutil.WriteFile(dstfile, data, 0x644)
	if wterr != nil {
		panic(err.Error())
	}
}
