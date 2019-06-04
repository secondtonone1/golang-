package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := strings.Replace("abcacmacd", "ac", "new", 2)
	fmt.Printf("str1 is : %s\n", str1)

	count := strings.Count(str1, "new")
	fmt.Printf("new in %s is %d counts\n", str1, count)

	str2 := strings.Repeat("abc", 3)
	fmt.Printf("str2 is : %s\n", str2)

	str3 := strings.ToLower("Nice to meet u, Bob")
	fmt.Printf("str3 is : %s\n", str3)

	str4 := strings.ToUpper("We have already got awards")
	fmt.Printf("str4 is : %s\n", str4)

	str5 := strings.TrimSpace("  abd  ")
	fmt.Printf("str5 is : %s\n", str5)

	str6 := strings.Trim("nofilmsno", "no")
	fmt.Printf("str6 is %s\n", str6)

	str7 := strings.TrimLeft("nofilmsno", "no")
	fmt.Printf("str7 is %s\n", str7)

	str8 := strings.TrimRight("nofilmsno", "no")
	fmt.Printf("str8 is %s\n", str8)

	strslice1 := strings.Split("What matters is not the good we expected others to do , but the good we do ", " ")
	fmt.Printf("strslice1 is : %v\n", strslice1)

	strslice2 := strings.Fields("What matters is not the good we expected others to do , but the good we do ")
	fmt.Printf("strslice2 is : %v\n", strslice2)

	strslice3 := []string{"Hello", "World", "!"}
	str9 := strings.Join(strslice3, " ")
	fmt.Printf("str9 is : %s\n", str9)
	num1, err := strconv.Atoi("156")
	if err != nil {
		fmt.Println("strconv.Atoi error")
	} else {
		fmt.Printf("num is %d\n", num1)
	}

	str10 := strconv.Itoa(100)
	fmt.Printf("str10 is %s\n", str10)
}
