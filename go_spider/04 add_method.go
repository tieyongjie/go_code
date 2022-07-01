package main

import "fmt"

type MyInt int

func (m MyInt) SayHello() {
	fmt.Println("hello, 我是一个老六(int)")
}
func main() {
	var m1 MyInt
	m1.SayHello()
	m1 = 100
	// %#v 先输出结构体名字值, 再输出结构体(字段类型 + 字段值)
	fmt.Printf("%#v", m1)
}
