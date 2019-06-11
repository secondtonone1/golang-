package main

import "fmt"

func main() {

	//数组初始化
	var arr1 [5]int = [5]int{5, 4, 3}
	fmt.Println(arr1)
	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	var age2 = [...]int{2, 3, 4, 5}
	fmt.Println(age2)
	var strs = [5]string{3: "good", 4: "moring"}
	fmt.Println(strs)

	//2维数组
	var multiarr [2][3]int
	fmt.Println(multiarr)
	var multiarr1 [2][3]int = [...][3]int{{3, 4, 5}, {6, 7, 8}}
	fmt.Println(multiarr1)
	for i := 0; i < len(multiarr1); i++ {
		for j := 0; j < len(multiarr1[i]); j++ {
			fmt.Println(multiarr1[i][j])
		}
	}

	for i1, v1 := range multiarr1 {
		for i2, v2 := range v1 {
			fmt.Println(i1, i2, v2)
		}
	}

}
