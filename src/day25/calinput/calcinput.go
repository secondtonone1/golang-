package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please input a string")
	var strinput string
	var err error
	var inputReader *bufio.Reader = bufio.NewReader(os.Stdin)
	strinput, err = inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("Your input is error")
	}
	fmt.Println("Your input is ", strinput)
	n_num := 0
	n_enchar := 0
	n_else := 0
	n_space := 0
	for _, v := range strinput {
		if v >= '0' && v <= '9' {
			n_num++
			continue
		}
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			n_enchar++
			continue
		}
		if v == ' ' {
			n_space++
			continue
		}
		n_else++
	}

	fmt.Println("空格字符个数", n_space)
	fmt.Println("数字个数", n_num)
	fmt.Println("英文字符个数", n_enchar)
	fmt.Println("其他字符的个数", n_else)
}
