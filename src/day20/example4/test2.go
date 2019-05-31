package main

import (
	"fmt"
	"strings"
)

func AddPrefix(inputStr *string) {
	if !strings.HasPrefix(*inputStr, "http://") {
		*inputStr = ("http://" + *inputStr)
	}
}

func AddSuffix(inputStr *string) {
	if !strings.HasSuffix(*inputStr, "/") {
		*inputStr = (*inputStr + "/")
	}
}

func FindStr(inputStr string, strFind string) {

	index := strings.Index(inputStr, strFind)
	if index == -1 {
		fmt.Printf("didn't find %s in %s \n", strFind, inputStr)
		return
	}

	fmt.Printf("found %s in %s, index is %d \n", strFind, inputStr, index)
}

func FindLastStr(inputStr string, strFind string) {
	index := strings.LastIndex(inputStr, strFind)
	if index == -1 {
		fmt.Printf("didn't find %s in %s \n ", strFind, inputStr)
		return
	}

	fmt.Printf("found %s in %s, index is %d \n", strFind, inputStr, index)
}

func main() {
	var inputStr string = ""
	fmt.Printf("input a string please:  ")
	fmt.Scanf("%s\n", &inputStr)
	AddPrefix(&inputStr)
	fmt.Printf("after add prefix input str is : %s\n", inputStr)

	fmt.Printf("input a string please:  ")
	fmt.Scanf("%s\n", &inputStr)
	AddSuffix(&inputStr)
	fmt.Printf("after AddSuffix input str is : %s\n", inputStr)

	FindStr("abcdefg", "mma")
	FindStr("abcabcefg", "bc")
	FindLastStr("abcdefg", "mma")
	FindLastStr("abcabcefg", "bc")
}
