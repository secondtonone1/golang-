package main

import (
	"flag"
	"fmt"
)

func main() {
	var bvar bool
	var nvar int
	var strvar string
	flag.BoolVar(&bvar, "b", false, "bool variable")
	flag.IntVar(&nvar, "n", 1024, "num variable")
	flag.StringVar(&strvar, "str", "Hello World", "string variable")
	flag.Parse()
	fmt.Println(bvar)
	fmt.Println(nvar)
	fmt.Println(strvar)
}
