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

//不定参数
func myfuncv(args ...int) {
    for _, arg := range args{
        fmt.Println(arg)
    }
}

//不定参数2
func myfuncv2(args []int){
    for _, arg := range args{
        fmt.Println(arg)
    }
}

func myfuncv3(args ...int){
    //按原样传递
    myfuncv(args...)
    //按切片传递
    myfuncv(args[1:]...)
}

func MyPrintf(args ...interface{}){
    for _, arg := range args{
        //interface 任意类型，
        //arg.(type)只能用于switch结构
        switch arg.(type){
        case int:
            fmt.Println(arg,"is an int value.")
        case string:
            fmt.Println(arg,"is a string value.")
        case int64:
            fmt.Println(arg,"is an int64 value")
        default:
            fmt.Println(arg,"is an unknown type")
        }
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
   
   myfuncv(2,3,5,6) 
   myfuncv2([]int{1,3,7,12})
   myfuncv3(21,32,15);

   var va1 int = 1
   var va2 int64 = 123
   var va3 string = "hello"
   var va4 float32 = 1.345
   MyPrintf(va1, va2, va3, va4)

   var cvar int = 5
   cfunc := func()(func()){
        return func(){
            fmt.Printf("cvar*2 : %d\n",cvar*2)
        }
   }()

   cfunc()
   cvar *= 2
   cfunc()
   
}