package main

import "fmt"

func main() {
	var a *int
	fmt.Println(a)
	var b int = 4
	a = &b
	//a = b
	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println(*&a)
	*a = 5 // 改变a的地址的值
	fmt.Println(a)
	fmt.Println(b) // b改变,因为a的地址是指向b的
}
