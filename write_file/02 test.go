package main

import "github.com/wutiange/go/sliceHandle"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	a, _ = sliceHandle.AddOrDelS(a, 1, 1)
}
