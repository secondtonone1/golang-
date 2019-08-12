package main

import "fmt"

func RevertString2(str string) string {
	lens := len(str)
	var strrt string
	for i := 0; i < lens; i++ {
		strrt += fmt.Sprintf("%c", str[lens-1-i])
	}
	return strrt
}

func RevertString(str string) string {
	var bytedata []byte
	bstr := []byte(str)
	lens := len(bstr)
	for i := 0; i < lens; i++ {
		bytedata = append(bytedata, bstr[lens-1-i])
	}
	return string(bytedata)
}

func main() {
	var str1 = "nice"
	var str2 = "to meet"
	var str3 = "u"
	fmt.Printf("%s\n", str1+" "+str2+" "+str3)
	str4 := fmt.Sprintf("%s %s %s\n", str1, str2, str3)
	fmt.Printf("%s\n", str4)
	fmt.Printf("str4's length is %d\n", len(str4))

	substr := str4[0:4]
	fmt.Println(substr)
	substr = str4[5:12]
	fmt.Println(substr)
	substr = str4[13:]
	fmt.Println(substr)

	fmt.Println(RevertString2("Hello World!"))
	fmt.Println(RevertString("Hello World!"))
}
