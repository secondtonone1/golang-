package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile, srcerr := os.Open("helloworld.txt")
	if srcerr != nil {
		fmt.Println("open src File failed")
		return
	}
	defer srcFile.Close()
	dstFile, dsterr := os.OpenFile("hellocopy.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if dsterr != nil {
		return
	}
	defer dstFile.Close()
	io.Copy(dstFile, srcFile)
}
