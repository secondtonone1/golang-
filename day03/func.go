package main
import "fmt"

func modify(array [5]int){
    array[0] = 200
    fmt.Println("In modify(), array values:", array)
}

func main(){
    //数组声明方法
    var bytearray [8]byte  //长度为8的数组
    fmt.Println(bytearray)
    var pointarray [4]* float64 //指针数组
    fmt.Println(pointarray)
    var mularray [3][5]int
    fmt.Println(mularray)
    fmt.Printf(" pointarray len is %v\n", len(pointarray))
    //数组遍历
    for i := 0; i < len(pointarray); i++{
        fmt.Println("Element", i, "of array is", pointarray[i])
    }

    //采用range遍历
    for i,v := range pointarray{
        fmt.Println("Array element [",i,"]=",v)
    }

    array := [5]int{1,2,3,4,5}
    modify(array)
    fmt.Println("In main(), array values:", array)
    //数组切片
    var mySlice []int = array[:3]
    fmt.Println("Elements of array")
    for _, v := range array{
        fmt.Print(v," ")
    }
    fmt.Println("\nElements of mySlice: ")
    for _, v := range mySlice{
        fmt.Print(v, " ")
    }

    fmt.Println()
    //直接创建元素个数为5的数组切片
    mkslice := make([]int ,5)
    fmt.Println("\n",mkslice)
    //创建初始元素个数为5的切片，元素都为0，且预留10个元素存储空间
    mkslice2 := make([]int,5,10)
    fmt.Println("\n",mkslice2)
    mkslice3 := []int{1,2,3,4,5}
    fmt.Println("\n",mkslice3)

    //元素遍历
    for i := 0; i <len(mkslice3); i++{
        fmt.Println("mkslice3[",i,"] =", mkslice3[i])
    }

    //range 遍历
    for i, v:= range mkslice3{
        fmt.Println("mkslice3[",i,"] =",v)
    }
    //获取size和capacity
    mkslice4 := make([]int,5,10)
    fmt.Println("len(mkslice4):", len(mkslice4))
    fmt.Println("cap(mkslice4):", cap(mkslice4))

    //末尾添加三个元素
    mkslice4 = append(mkslice4,1,2,3)
    fmt.Println("mkslice4 is : ", mkslice4)

    mkslice4 = append(mkslice4, mkslice3...)
    fmt.Println("mkslice4 is : ", mkslice4)

    oldslice := []int{1,2,3,4,5}
    newslice := oldslice[:3]
    fmt.Println("newslice is :", newslice)
}