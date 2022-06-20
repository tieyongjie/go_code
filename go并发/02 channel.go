package main

import "fmt"

/*
通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，
发送或接收。如果未指定方向，则为双向通道。
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum /*把sum发送给通道c */
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	x := <-c
	fmt.Println(x)
	go sum(s[len(s)/2:], c)
	y := <-c
	fmt.Println(x, y, x+y)
}
