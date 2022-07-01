package main

import "fmt"

type Student struct {
	Id      int64
	Name    string
	Class   string
	Age     int64
	Address string
}

func main() {
	Students := make([]Student, 0)
	fmt.Println(Students)
}
