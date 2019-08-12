package main

import(
	"fmt"
	"strings"
)

func makeFuncAddSuffix(suffixstr string)func(string)string{
	return func(name string) string{
		if ( !strings.HasSuffix(name, suffixstr)){
			return name + suffixstr
		}
		return name
	}
}

func main() {
	funcjpg:=makeFuncAddSuffix(".jpg")
	funcbmp:=makeFuncAddSuffix(".bmp")
	fmt.Println(funcjpg("name"))
	fmt.Println(funcbmp("ware"))
}
