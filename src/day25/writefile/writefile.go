package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputErr := os.OpenFile("helloworld.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputErr != nil {
		fmt.Println("openfile error")
		return
	}

	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "Hello Hurricane!\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
