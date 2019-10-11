package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
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

type Hero struct {
	Name    string
	Attack  int
	Defence int
	GenTime int64
}

type HeroList []*Hero

func (hl HeroList) Len() int {
	return len(hl)
}

func (hl HeroList) Less(i, j int) bool {
	if i < 0 || j < 0 {
		return true
	}

	lenth := len(hl)
	if i >= lenth || j >= lenth {
		return true
	}

	if hl[i].Attack != hl[j].Attack {
		return hl[i].Attack < hl[j].Attack
	}

	if hl[i].Defence != hl[j].Defence {
		return hl[i].Defence < hl[j].Defence
	}

	return hl[i].GenTime < hl[j].GenTime
}

func (hl HeroList) Swap(i, j int) {
	if i < 0 || j < 0 {
		return
	}

	lenth := len(hl)
	if i >= lenth || j >= lenth {
		return
	}

	hl[i], hl[j] = hl[j], hl[i]

}

func main() {
	arrayint := []int{6, 1, 0, 5, 2, 7}
	sort.Ints(arrayint)
	fmt.Println(arrayint)
	arraystring := []string{"hello", "world", "Alis", "and", "Bob"}
	sort.Strings(arraystring)
	fmt.Println(arraystring)
	//自定义类型排序用sort.Sort
	var herolists HeroList
	for i := 0; i < 10; i++ {
		generate := time.Now().Unix()
		name := fmt.Sprintf("Hero%d", generate)
		hero := Hero{
			Name:    name,
			Attack:  rand.Intn(100),
			Defence: rand.Intn(200),
			GenTime: generate,
		}
		herolists = append(herolists, &hero)
		time.Sleep(time.Duration(1) * time.Second)
	}

	sort.Sort(herolists)
	for _, value := range herolists {
		fmt.Print(value.Name, " ", value.Attack, " ", value.Defence, " ", value.GenTime, "\n")
	}
}
