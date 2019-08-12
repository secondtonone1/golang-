package main

import (
	"fmt"
)

func main() {
	//label写在外层循环上边。continue时，跳出内层循环，会继续执行外层循环
LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is %d, j is %d \n", i, j)
		}
	}

	fmt.Println("*********************")

	//label写在内层循环上边。continue时，跳出内层循环，会继续执行外层循环
	//和写在外层循环一个效果
	for i := 0; i < 5; i++ {
	LABEL3:
		for j := 0; j < 5; j++ {
			if j == 4 {
				continue LABEL3
			}
			fmt.Printf("i is %d, j is %d \n", i, j)
		}
	}

	fmt.Println("++++++++++++++++++++++++++")
	//break 标签后，跳出内层循环，继续执行外层循环。
	for i := 0; i < 5; i++ {
	LABEL2:
		for j := 0; j < 5; j++ {
			if j == 4 {
				break LABEL2
			}
			fmt.Printf("i is %d, j is %d \n", i, j)
		}
	}

	fmt.Println("+++++++++++++++++++")
	//break 标签后，跳出两层循环，不再执行外层和内层循环。
LABEL4:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 4 {
				break LABEL4
			}
			fmt.Printf("i is %d, j is %d \n", i, j)
		}
	}
	//goto到外层标签,会继续执行循环，从而不断循环

	/*
	   LABEL5:
	   	for i := 0; i < 5; i++ {
	   		for j := 0; j < 5; j++ {
	   			if j == 4 {
	   				goto LABEL5
	   			}
	   			fmt.Printf("i is %d, j is %d \n", i, j)
	   		}
	   	}
	*/

	inex1 := 0
HERE:
	fmt.Println("index is", inex1)
	inex1++
	if inex1 == 5 {
		return
	}

	goto HERE
}
