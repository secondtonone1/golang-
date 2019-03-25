package main
import "fmt"
import "errors"
import "mymaths"

func Add(a int, b int) (ret int, err error){
    if a < 0 || b < 0{
        //不允许负数相加
        err= errors.New("should be non-negative numbers")
        return 
    }
    return a + b, nil
}


 //goto跳转到函数某个标签
 func myfunc(){
    i := 0
    HERE:
    fmt.Println(i)
    i++
    if i < 10 {
        goto HERE
    }
}

func main(){
    //控制语句
    var a int = 5
    if (a <10){
        fmt.Println("a < 10")
    }else{
        fmt.Println("a >= 10")
    }
    //选择语句
    i:=2
    switch i{
    case 0:
        fmt.Println("0")
    case 1:
        fmt.Printf("1")
    case 2:
        fallthrough
    case 3:
        fmt.Printf("3")
    case 4, 5, 6:
        fmt.Printf("4, 5, 6")
    default:
        fmt.Printf("Default")
    }

    //循环
    sum := 0
    for i :=0; i < 10; i++{
        sum += i
    }

    sum1 := 0
    for {
        sum1 ++
        if sum1 > 10{
            break
        }
    }
    fmt.Printf("\n")
    array := []int{1,2,3,4,5,6}
    //数组首尾交换
    for i, j := 0, len(array) -1; i <j ; i, j =i+1, j-1{
        array[i], array[j] = array[j], array[i]
    }
    
    for index, value :=  range array{
        fmt.Println("index is", index)
        fmt.Println("value is", value)
    }

   myfunc()

   addres, errres := mymaths.Add(1,2)
   fmt.Println("addres: ",addres, "errres: ", errres)

}