package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("helloworld.txt")
	if err != nil {
		fmt.Println("open file failed!")
		return
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerr := inputReader.ReadString('\n')
		if readerr == io.EOF {
			fmt.Println("Read file end")
			return
		}

		if readerr != nil {
			fmt.Println("Read string except")
			return
		}

		fmt.Printf("Has read string is %s", inputString)
	}
}
