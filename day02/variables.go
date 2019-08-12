package main
import "fmt"
var v1 int  //整形
var v2 string //字符串
var v3 [10]int  //数组
var v4 []int   //切片
var v5 struct{   //结构体
    f int
}
var v6 * int  //指针
var v7 map[string]int  //map，key为string,value为int
var v8 func(a int) int  //函数对象

var(
    v9 int
    v10 string
)

//匿名变量
func GetName()(firstName, lastName, nickName string){
    return "first","last","nick"
}

func main(){
    //变量赋值
    v1 = 10
    v2 := "hello"
    fmt.Println("v2:",v2)
    //变量初始化
    v11  := "day02"
    fmt.Println("v11:",v11)
    //第二种初始化方式
    var v12 = 13
    //变量交换
    v1, v12 = v12, v1
    fmt.Printf("v11:%v, v12:%v\n",v1,v12)
    _, _, nickName := GetName()
    fmt.Printf("nickName : %v\n", nickName)
    //常量
    const Pi float64 = 3.141592653
    const zero = 0.0
    const (
        size int64 = 1024
        eof = -1
    )
    const u, v float32 = 0,3
    const a, b, c = 3, 4, "foo"
    //iota 表示初始化常量为0，之后每次出现iota，iota自增1
    const ( // iota被重设为0
        c0 = iota // c0 == 0
        c1 = iota // c1 == 1
        c2 = iota // c2 == 2
        )
    
    const (
        a1 = 1 << iota  //1左移0位
        b2 = 1 << iota  //1左移1位
        c3 = 1 << iota  //1左移2位
    )
    fmt.Printf("a1 : %v, b2 : %v, c3 : %v\n", a1, b2, c3)

    const (
        Sunday = iota
        Monday
        Tuesday
        Wednesday
        Thursday
        Friday
        Saturday
        numberOfDays // 这个常量没有导出
        )
        //同Go语言的其他符号（symbol）一样，以大写字母开头的常量在包外可见

    //bool 类型
    var bvar bool 
    bvar = true
    bvar2 := (1==2) //bvar 被推导为bool类型
    fmt.Printf("bvar: %v, bvar2: %v\n",bvar, bvar2)
    
    //字符串
    var str string
    str = "Hello world"
    ch := str[0]
    fmt.Printf("The length of \"%s\" is %d \n ", str, len(str))
    fmt.Printf("The first character of \"%s\" is %c.\n ",str, ch)
    str2 := ",I'm mooncoder"
    fmt.Println(str+str2)

    hellos := "Hello,我是中国人"
    //字符串遍历,utf-8编码遍历,每个字符byte类型，8字节
    for i:=0; i < len(hellos); i++{
        fmt.Printf("%c", hellos[i])
   }
   //unicode 遍历,每个字符rune类型，变长
   for i, ch := range hellos{
       fmt.Printf("%c\t",ch)
       fmt.Printf("%d\n",i)
   }
   
}