package main

import (
	"fmt"
	"sort"
)

func main() {

	//数组初始化
	var arr1 [5]int = [5]int{5, 4, 3}
	var slice1 []int = arr1[1:len(arr1)]
	fmt.Println(slice1)
	slice2 := []int{1, 3, 5}
	var slice3 []int = make([]int, 3, 6)
	slice3 = append(slice3, slice1...)
	fmt.Println(slice3)
	slice3 = append(slice3, slice2...)
	fmt.Println(slice3)

	str := "hello world"
	s := []byte(str)
	s[0] = 'g'
	str = string(s)
	fmt.Println(str)

	str2 := "我爱中国"
	s2 := []rune(str2)
	s2[0] = '你'
	str2 = string(s2)
	fmt.Println(str2)
	//切片排序
	sort.Ints(slice3)
	fmt.Println(slice3)
	fmt.Println(sort.SearchInts(slice3, 3))

}
