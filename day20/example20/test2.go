package main

import (
	"fmt"
	"sort"
)

func modify(data map[string]int, key string, value int) {
	v, res := data[key]
	//不存在,则res是false
	if !res {
		fmt.Println("key not find")
		return
	}
	fmt.Println("key is ", key, "value is ", v)
	data[key] = value
}

func sortprintmap(data map[string]int) {
	slice := make([]string, 0)
	for k, _ := range data {
		slice = append(slice, k)
	}
	sort.Strings(slice)
	for _, s := range slice {
		d, e := data[s]
		if !e {
			continue
		}
		fmt.Println("key is ", s, "value is ", d)

	}
}

func main() {
	var data map[string]int = map[string]int{"bob": 18, "luce": 28}
	modify(data, "lilei", 28)
	modify(data, "luce", 22)
	fmt.Println(data)

	//map大小
	fmt.Println(len(data))
	//map 使用前一定要初始化，可以显示初始化，也可以用make
	var data2 map[string]int = make(map[string]int, 3)
	fmt.Println(data2)
	//当key不存在时，则会插入
	data2["sven"] = 19
	fmt.Println(data2)
	//当key存在时，则修改
	data2["sven"] = 299
	fmt.Println(data2)
	data2["Arean"] = 33
	data2["bob"] = 178
	//map是无序的,遍历输出
	for key, value := range data2 {
		fmt.Println("key: ", key, "value: ", value)
	}
	//可以将key排序然后输出即可
	sortprintmap(data2)

	//二维map
	var usrdata map[string]map[string]int
	//使用前需要初始化
	usrdata = make(map[string]map[string]int)
	usrdata["sven"] = make(map[string]int)
	usrdata["sven"]["age"] = 21
	usrdata["sven"]["id"] = 1024

	usrdata["susan"] = make(map[string]int)
	usrdata["susan"]["age"] = 19
	usrdata["susan"]["id"] = 1000

	//二维map 遍历
	for k, v := range usrdata {
		for k2, v2 := range v {
			fmt.Println(k, " ", k2, " ", v2)
		}
	}

	//slice of map
	slicem := make([]map[string]int, 5)
	for i := 0; i < len(slicem); i++ {
		slicem[i] = make(map[string]int)
	}

	fmt.Println(slicem)

	//map 反转
	rvmap := make(map[int]string)
	for k, v := range data2 {
		rvmap[v] = k
	}
	fmt.Println(rvmap)
	fmt.Println(data2)
}
