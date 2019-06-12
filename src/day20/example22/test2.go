package main

import "fmt"

func quickSort(slice []int, begin int, end int) {
	if len(slice) == 0 {
		return
	}

	if end >= len(slice) {
		return
	}

	if begin < 0 {
		return
	}

	if begin >= end {
		return
	}

	i := begin
	j := end
	value := slice[i]

	for {
		index := j
		for ; index > i; index-- {
			if value > slice[index] {
				slice[i] = slice[index]
				//slice[index] = value
				break
			}
		}
		j = index
		for index = i; index < j; index++ {
			if value < slice[index] {
				slice[j] = slice[index]
				//slice[index] = value
				break
			}
		}
		i = index

		//fmt.Println(i, j)
		if i >= j {
			slice[i] = value
			quickSort(slice, begin, i-1)
			quickSort(slice, i+1, end)
			return
		}

	}

}

func bubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[i], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func selectSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

func insertSort(slice []int) {
	if len(slice) <= 0 {
		return
	}

	for i := 1; i < len(slice); i++ {
		temp := slice[i]
		for j := i - 1; j >= 0; j-- {
			if temp < slice[j] {
				slice[j+1] = slice[j]
			}
			slice[j+1] = temp
			break
		}
	}
}

func main() {
	array := [...]int{6, 2, 7, 3, 8, 9}
	slice := array[:]
	quickSort(slice, 0, len(slice)-1)
	fmt.Println(slice)
	slice2 := array[:]
	bubbleSort(slice2)
	fmt.Println(slice2)
	slice3 := array[:]
	selectSort(slice3)
	fmt.Println(slice3)
	slice4 := array[:]
	insertSort(slice4)
	fmt.Println(slice4)

}
