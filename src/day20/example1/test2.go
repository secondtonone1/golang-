package main

import "fmt"

func main() {
	var n int
	var flags bool = true
	fmt.Scanf("%d", &n)

	for i := 2; i < n; i++ {
		if n%i == 0 {
			flags = false
			break
		}
	}

	if flags == true {
		fmt.Println(n, " is su shu")
	} else {
		fmt.Println(n, "is not sushu")
	}
}
