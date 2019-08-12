package main

import "fmt"

type BigRootHeap struct {
	m_array []int
}

func (bh *BigRootHeap) swap(array []int, i, j int) {
	
	if i >= len(array) || j >= len(array) {
		return
	}

	temp := array[i]
	array[i] = array[j]
	array[j] = temp

}

func (bh *BigRootHeap) initHeap(array []int) {
	bh.m_array = make([]int, len(array))
	copy(bh.m_array, array)
	length := len(bh.m_array)

	//每次循环后长度减少，因为每次循环最后元素都变为最大
	for index := length/2 - 1; index >= 0; index-- {
		bh.adjustDown(bh.m_array, index, length)
	}
}

func (bh *BigRootHeap) adjustDown(array []int, index, length int) {
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

		//如果堆顶元素最大，则退出
		if array[index] > array[maxchild] {
			break
		}

		//堆顶元素小于其子节点，则交换，并且将堆顶元素继续下调
		bh.swap(array, index, maxchild)
		index = maxchild
		leftchild = index*2 + 1
		rightchild = leftchild + 1
		maxchild = leftchild

	}
}

func (bh *BigRootHeap) popBigRoot() (bool, int) {
	length := len(bh.m_array)
	if length <= 0 {
		return false, 0
	}
	temp := bh.m_array[0]
	bh.swap(bh.m_array, 0, length-1)
	bh.m_array = append(bh.m_array[0:length-1], bh.m_array[length:]...)
	bh.adjustDown(bh.m_array, 0, len(bh.m_array))
	return true, temp
}

func (bh* BigRootHeap) adjustUp(array []int, index, length int){
	
	parent := (index+1)/2-1
	for{
		if index <= 0{
			break
		}

		if(bh.m_array[index] <= bh.m_array[parent]){
			break
		}

		bh.swap(bh.m_array,index,parent)
		index = parent
		parent = (index+1)/2-1

	}
}

func (bh* BigRootHeap) insertNode(node int){
	bh.m_array = append(bh.m_array,node)
	length:=len(bh.m_array)
	index := length-1
	bh.adjustUp(bh.m_array,index,length)

}

func main() {

	//数组初始化
	array := []int{7, 1, 0, 5, 2, 9, 6}
	bh := new(BigRootHeap)
	bh.initHeap(array)
	for{
		res,num:=bh.popBigRoot()
		if(!res){
			break
		}
		fmt.Println(num)
	}
	fmt.Println("")
	bh.initHeap(array)
	bh.insertNode(10)
	bh.insertNode(3)
	for{
		res,num:=bh.popBigRoot()
		if(!res){
			break
		}
		fmt.Println(num)
	}
}
