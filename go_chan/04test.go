package main

import "fmt"

func main() {
	c := make(chan int, 10)
	fmt.Println(cap(c))
}
