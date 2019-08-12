package main

import (
	"bufio"
	"fmt"
	"strings"
)

//bufio实现带缓存的I/O操作
//buffio 基于io.Reader 或io.Writer 封装为Reader类型
/*
type Reader struct {
	buf          []byte
	rd           io.Reader // reader provided by the client
	r, w         int       // buf read and write positions
	err          error
	lastByte     int
	lastRuneSize int
}
*/

func main() {

	// NewReaderSize 将 rd 封装成一个拥有 size 大小缓存的 bufio.Reader 对象
	// 如果 rd 的基类型就是 bufio.Reader 类型，而且拥有足够的缓存
	// 则直接将 rd 转换为基类型并返回
	//func NewReaderSize(rd io.Reader, size int) *Reader
	// NewReader 相当于 NewReaderSize(rd, 4096)
	//func NewReader(rd io.Reader) *Reader

	sreader := strings.NewReader("Hello world!")
	bufreader := bufio.NewReader(sreader)

	// Peek 返回缓存的一个切片，该切片引用缓存中前 n 字节数据
	// 该操作不会将数据读出，只是引用
	// 引用的数据在下一次读取操作之前是有效的
	//Peek传递的数值应小于bufio的缓存大小

	b, _ := bufreader.Peek(5)
	fmt.Printf("%s\n", b)

	b[0] = 'a'
	b, _ = bufreader.Peek(5)
	fmt.Printf("%s\n", b)

	//从bufio中读取数据到[]byte中，返回读取的字节数
	//如果缓冲区有数据，则优先读取缓冲区数据
	//如果缓冲区没有数据，则读取bufio.Reader中的数据

	s := strings.NewReader("Hello World!")
	br := bufio.NewReader(s)
	bytes := make([]byte, 6)
	//先读取6个
	n, err := br.Read(bytes)
	fmt.Printf("%-6s %-2v %v\n", bytes[:n], n, err)
	//在读取6个
	n, err = br.Read(bytes)
	fmt.Printf("%-6s %-2v %v\n", bytes[:n], n, err)
	//bufio没有数据了，返回EOF错误
	n, err = br.Read(bytes)
	fmt.Printf("%-6s %-2v %v\n", bytes[:n], n, err)

	//读取一个字节
	s2 := strings.NewReader("Hello World!")
	br2 := bufio.NewReader(s2)

	c, _ := br2.ReadByte()
	fmt.Printf("%c\n", c)

	c, _ = br2.ReadByte()
	fmt.Printf("%c\n", c)
	//撤回读出的数据
	br2.UnreadByte()
	c, _ = br2.ReadByte()
	fmt.Printf("%c\n", c)
	// ReadRune 从 bufio中读出一个 UTF8 编码的字符并返回
	// 同时返回该字符的 UTF8 编码长度
	strio := strings.NewReader("你好，我是码农")
	chbr := bufio.NewReader(strio)

	rc, size, _ := chbr.ReadRune()
	fmt.Printf("%c %v\n", rc, size)

	rc, size, _ = chbr.ReadRune()
	fmt.Printf("%c %v\n", rc, size)

	br.UnreadRune()
	rc, size, _ = chbr.ReadRune()
	fmt.Printf("%c %v\n", rc, size)

}
