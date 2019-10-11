package main

import (
	"fmt"
	"math/rand"
)

/*
// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
   }
*/
//接口起到规范作用，实现参数类型的限制

type LinkList struct {
	Head *LinkEle
	Tail *LinkEle
}

type LinkEle struct {
	Data interface{}
	Pre  *LinkEle
	Next *LinkEle
}

func (le *LinkEle) GetData() interface{} {
	return le.Data
}

func (ll *LinkList) InsertTail(le *LinkEle) {
	if ll.Tail == nil && ll.Head == nil {
		ll.Tail = le
		ll.Head = ll.Tail
		return
	}
	ll.Tail.Next = le
	le.Pre = ll.Tail
	le.Next = nil
	ll.Tail = le
}

func (ll *LinkList) InsertHead(le *LinkEle) {

	if ll.Tail == nil && ll.Head == nil {
		ll.Tail = le
		ll.Head = ll.Tail
		return
	}

	ll.Head.Pre = le
	le.Pre = nil
	le.Next = ll.Head
	ll.Head = le
}

func (ll *LinkList) InsertIndex(le *LinkEle, index int) {
	if index < 0 {
		return
	}

	if ll.Head == nil {
		ll.Head = le
		ll.Tail = ll.Head
		return
	}

	node := ll.Head
	indexfind := 0
	for ; indexfind < index; indexfind++ {
		if node.Next == nil {
			break
		}
		node = node.Next
	}

	if indexfind != index {
		fmt.Println("index is out of range")
		return
	}
	//node 后边的节点缓存起来
	nextnode := node.Next

	//node 和le连接起来
	node.Next = le
	le.Pre = node

	if node == ll.Tail {
		ll.Tail = le
		return
	}

	//le和next node 连接起来
	if nextnode != nil {
		nextnode.Pre = le
		le.Next = nextnode
	}
}

func (ll *LinkList) DelIndex(index int) {
	if index < 0 {
		return
	}

	if ll.Head == nil {
		return
	}

	node := ll.Head
	indexfind := 0
	for ; indexfind < index; indexfind++ {
		if node.Next == nil {
			break
		}
		node = node.Next
	}

	if indexfind != index {
		fmt.Println("index is out of range")
		return
	}

	if ll.Head == ll.Tail {
		ll.Tail = nil
		ll.Head = ll.Tail
		return
	}

	//如果是头节点
	if node == ll.Head {
		ll.Head = node.Next
		node.Next.Pre = nil
		return
	}

	//如果是尾结点
	if node == ll.Tail {
		ll.Tail = node.Pre
		ll.Tail.Next = nil
		return
	}
	//将前后连接起来
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (ll *LinkList) DelHead() {
	if ll.Head == ll.Tail && ll.Head == nil {
		return
	}

	if ll.Head == ll.Tail {
		ll.Head = nil
		ll.Tail = ll.Head
		return
	}

	ll.Head = ll.Head.Next
	ll.Head.Pre = nil
}

func (ll *LinkList) DelTail() {
	if ll.Head == ll.Tail && ll.Head == nil {
		return
	}

	if ll.Head == ll.Tail {
		ll.Head = nil
		ll.Tail = ll.Head
		return
	}

	ll.Tail = ll.Tail.Pre
	ll.Tail.Next = nil
}

func main() {
	//rand.Seed(time.Now().Unix())
	ll := &LinkList{nil, nil}
	fmt.Println("insert head .....................")
	for i := 0; i < 2; i++ {
		num := rand.Intn(100)
		node1 := &LinkEle{Data: num, Next: nil, Pre: nil}
		ll.InsertHead(node1)
		fmt.Println(num)
	}
	fmt.Println("after insert head .................")
	for node := ll.Head; node != nil; node = node.Next {
		val, ok := node.GetData().(int)
		if !ok {
			fmt.Println("interface transfer error")
			break
		}
		fmt.Println(val)
	}

	fmt.Println("insert tail .....................")
	for i := 0; i < 2; i++ {
		num := rand.Intn(100)
		node1 := &LinkEle{Data: num, Next: nil, Pre: nil}
		ll.InsertTail(node1)
		fmt.Println(num)
	}

	fmt.Println("after insert tail .................")
	for node := ll.Head; node != nil; node = node.Next {
		val, ok := node.GetData().(int)
		if !ok {
			fmt.Println("interface transfer error")
			break
		}
		fmt.Println(val)
	}

	fmt.Println("insert after third element........")
	{
		num := rand.Intn(100)
		node1 := &LinkEle{Data: num, Next: nil, Pre: nil}
		ll.InsertIndex(node1, 2)
		fmt.Println(num)
	}

	fmt.Println("after insert index .................")
	for node := ll.Head; node != nil; node = node.Next {
		val, ok := node.GetData().(int)
		if !ok {
			fmt.Println("interface transfer error")
			break
		}
		fmt.Println(val)
	}

	//delete second element, its index is 1
	fmt.Println("delete second element, its index is 1")
	ll.DelIndex(1)
	fmt.Println("after delete second element, its index is 1")
	for node := ll.Head; node != nil; node = node.Next {
		val, ok := node.GetData().(int)
		if !ok {
			fmt.Println("interface transfer error")
			break
		}
		fmt.Println(val)
	}

	fmt.Println("delete head")
	ll.DelHead()
	fmt.Println("after delete head")

	for node := ll.Head; node != nil; node = node.Next {
		val, ok := node.GetData().(int)
		if !ok {
			fmt.Println("interface transfer error")
			break
		}
		fmt.Println(val)
	}

	fmt.Println("delete tail")
	ll.DelTail()
	fmt.Println("after delete tail")

	for node := ll.Head; node != nil; node = node.Next {
		val, ok := node.GetData().(int)
		if !ok {
			fmt.Println("interface transfer error")
			break
		}
		fmt.Println(val)
	}
}
