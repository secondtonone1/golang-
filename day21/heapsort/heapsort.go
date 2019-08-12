package main

import "fmt"

type HeapSort struct {
}

//调整index为根的子树，此时index的左右子树都是大根树
//比较index和其左右节点，将index根节点设置为最大的元素
//可能会引起子树失效，所以会循环处理修改的子树
func (hs *HeapSort) adjustHeap(array []int, index, length int) {
	//index 的左右子节点
	leftchild := index*2 + 1
	rightchild := leftchild + 1
	maxchild := leftchild

	for {

		//如果左节点比长度大，说明该节点为子节点
		if leftchild > length-1 {
			break
		}
		//右节点存在，且比左节点大
		if rightchild <= length-1 && array[rightchild] > array[maxchild] {
			maxchild = rightchild
		}
		//index 元素比最大子节点大，则不需要交换，退出
		if array[index] > array[maxchild] {
			break
		}

		//比较自己元素和最大节点的元素，做交换
		hs.swap(array, index, maxchild)
		index = maxchild
		leftchild = index*2 + 1
		rightchild = leftchild + 1
		maxchild = leftchild
	}

}

func (hs *HeapSort) swap(array []int, i, j int) {
	if i >= len(array) || j >= len(array) {
		return
	}

	temp := array[i]
	array[i] = array[j]
	array[j] = temp

}

func (hs *HeapSort) sort(array []int) {

	//每次循环后长度减少，因为每次循环最后元素都变为最大
	for length := len(array); length > 1; length-- {
		//最后一个非叶子节点索引
		lastnode := length/2 - 1
		//从最后一个非叶子节点一次从左到右，从下到上排序子树
		//循环过后得到一个大顶堆(此时还不是大根堆)
		for i := lastnode; i >= 0; i-- {
			hs.adjustHeap(array, i, length)
		}

		//将堆顶元素放到末尾
		hs.swap(array, 0, length-1)
		fmt.Println(array)
	}

}

func main() {

	//数组初始化
	array := []int{6, 1, 0, 5, 2, 9, 6}
	hs := new(HeapSort)
	hs.sort(array)

}
