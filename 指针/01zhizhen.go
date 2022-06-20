package main

import "fmt"

func main() {
	var num int = 10
	//num内存地址，&num
	fmt.Printf("num address=%v\n", &num)
	var ptr = &num
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("*ptr=%v\n", *ptr)
}
